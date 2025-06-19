package cmd

import (
	"fmt"
	"os"

	"github.com/RickardAhlstedt/gitgeist/config"
	"github.com/RickardAhlstedt/gitgeist/detector"
	"github.com/RickardAhlstedt/gitgeist/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitgeist",
	Short: "Summon the Gitgeist to inspect your repo for sins",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("👻 The Gitgeist stirs...")

		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Println("🪦 Failed to load config:", err)
			os.Exit(1)
		}

		dir, err := internal.GetCurrentDir()
		if err != nil {
			fmt.Println("🪦 Failed to get current directory:", err)
			os.Exit(1)
		}
		fmt.Println("🔍 Inspecting repository at:", dir)

		msg, err := detector.GetLastCommitMessage()
		if err != nil {
			fmt.Println("🪦 Failed to get last commit message:", err)
			os.Exit(1)
		}
		fmt.Println("📜 Last commit message:", msg)
		warnings := detector.AnalyzeCommitMessage(msg, cfg.CommitMessagePatterns)
		if len(warnings) > 0 {
			fmt.Println("⚠️ Commit message warnings:")
			for _, warning := range warnings {
				fmt.Println("  -", warning)
			}
		} else {
			fmt.Println("✅ Commit message looks good!")
		}

		fmt.Println("\n🔎 Scanning repository files for suspicious patterns...")

		fileWarnings, err := detector.ScanRepoFiles(dir, cfg.FileInspectionPatterns)
		if err != nil {
			fmt.Println("Error scanning files:", err)
			return
		}

		if len(fileWarnings) > 0 {
			fmt.Println("⚠️ Gitgeist found suspicious patterns in files:")
			for _, warning := range fileWarnings {
				fmt.Println(" -", warning)
			}
		} else {
			fmt.Println("✅ No suspicious patterns found in repository files")
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("🪦 Gitgeist failed:", err)
		os.Exit(1)
	}
}
