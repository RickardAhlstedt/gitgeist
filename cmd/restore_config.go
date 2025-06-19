package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/RickardAhlstedt/gitgeist/config"
	"github.com/spf13/cobra"
)

var restoreConfigCmd = &cobra.Command{
	Use:   "restore-config",
	Short: "Restore the Gitgeist config by resetting ~/.gitgeist directory",
	Run: func(cmd *cobra.Command, args []string) {
		usr, err := user.Current()
		if err != nil {
			fmt.Println("Failed to get current user:", err)
			return
		}
		configDir := filepath.Join(usr.HomeDir, ".gitgeist")

		// Remove config directory if exists
		if _, err := os.Stat(configDir); err == nil {
			err = os.RemoveAll(configDir)
			if err != nil {
				fmt.Println("Failed to remove existing config directory:", err)
				return
			}
		}

		// Create config directory
		err = os.MkdirAll(configDir, 0755)
		if err != nil {
			fmt.Println("Failed to create config directory:", err)
			return
		}

		// Write default config.yaml
		defaultCfg := config.GetDefaultConfig()
		data, err := config.MarshalConfig(defaultCfg)
		if err != nil {
			fmt.Println("Failed to marshal default config:", err)
			return
		}

		configPath := filepath.Join(configDir, "config.yaml")
		err = os.WriteFile(configPath, data, 0644)
		if err != nil {
			fmt.Println("Failed to write default config file:", err)
			return
		}

		fmt.Println("✅ Gitgeist config restored at", configPath)
	},
}

func init() {
	rootCmd.AddCommand(restoreConfigCmd)
}
