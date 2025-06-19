package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type commit_message_length struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

type Config struct {
	CommitMessagePatterns  []string              `yaml:"commit_message_patterns"`
	CommitMessageLength    commit_message_length `yaml:"commit_message_length"`
	FileInspectionPatterns []string              `yaml:"file_inspection_patterns"`
	ProtectedBranches      []string              `yaml:"protected_branches"`
	GitAntiPatterns        []string              `yaml:"git_anti_patterns"`
}

func GetDefaultConfig() *Config {
	return &Config{
		CommitMessagePatterns: []string{
			`(?i)^fix$`,
			`(?i)^wip`,
			`(?i)final`,
			`(?i)temp`,
			`(?i)debug`,
		},
		CommitMessageLength: commit_message_length{
			Min: 10,
			Max: 100,
		},
		FileInspectionPatterns: []string{
			`console\.log`,
			`debugger`,
			`dd`,
			`dump`,
			`TODO`,
		},
	}
}

func LoadConfig() (*Config, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}
	configDir := filepath.Join(usr.HomeDir, ".gitgeist")
	configPath := filepath.Join(configDir, "config.yaml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Create directory if missing
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create config directory: %w", err)
		}
		// Write default config file
		defaultCfg := GetDefaultConfig()
		data, err := yaml.Marshal(defaultCfg)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal default config: %w", err)
		}
		if err := os.WriteFile(configPath, data, 0644); err != nil {
			return nil, fmt.Errorf("failed to write default config file: %w", err)
		}
		fmt.Printf("Created default config at %s\n", configPath)
	}

	// Load config file
	f, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	var cfg Config
	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}
	return &cfg, nil
}

func MarshalConfig(cfg *Config) ([]byte, error) {
	return yaml.Marshal(cfg)
}
