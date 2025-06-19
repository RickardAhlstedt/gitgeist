package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var installHookCmd = &cobra.Command{
	Use:   "install-hook",
	Short: "Install Gitgeist pre-commit hook",
	Long:  `This command installs the Gitgeist pre-commit hook to automatically analyze commit messages and repository files before each commit. It helps maintain code quality and catch potential issues early.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Logic to install the pre-commit hook
		err := installPreCommitHook()
		if err != nil {
			cmd.Println("🪦 Failed to install pre-commit hook:", err)
			return
		}
		cmd.Println("✅ Gitgeist pre-commit hook installed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(installHookCmd)
}

func installPreCommitHook() error {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("🪦 Failed to get current directory:", err)
		return nil
	}

	gitDir := ""
	for {
		check := filepath.Join(dir, ".git")
		info, err := os.Stat(check)
		if err == nil && info.IsDir() {
			gitDir = check
			break
		}

		parent := filepath.Dir(dir)
		if parent == dir { // reached root
			break
		}
		dir = parent
	}

	if gitDir == "" {
		return fmt.Errorf("not a git repository (no .git directory found)")
	}

	hooksDir := filepath.Join(gitDir, "hooks")
	if err := os.MkdirAll(hooksDir, 0755); err != nil {
		return fmt.Errorf("failed to create hooks directory: %w", err)
	}
	hookPath := filepath.Join(hooksDir, "pre-commit")
	hookContent := `#!/bin/sh
# Gitgeist pre-commit hook
gitgeist
if [ $? -ne 0 ]; then
  echo "Gitgeist checks failed. Commit aborted."
  exit 1
fi
`
	if err := os.WriteFile(hookPath, []byte(hookContent), 0755); err != nil {
		return fmt.Errorf("failed to write pre-commit hook: %w", err)
	}
	fmt.Println("🔧 Pre-commit hook installed at:", hookPath)
	return nil
}
