package detector

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/RickardAhlstedt/gitgeist/internal"
)

func ScanRepoFiles(root string, patterns []string) ([]string, error) {
	var results []string

	// Compile regex patterns once
	var regexes []*regexp.Regexp
	for _, p := range patterns {
		re, err := regexp.Compile(p)
		if err != nil {
			return nil, fmt.Errorf("invalid regex pattern '%s': %w", p, err)
		}
		regexes = append(regexes, re)
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Ignore permission errors, etc.
			return nil
		}
		// Skip directories we don't want to scan
		if info.IsDir() {
			dirname := info.Name()
			if dirname == ".git" || dirname == "node_modules" || strings.HasPrefix(dirname, ".") {
				return filepath.SkipDir
			}
			return nil
		}

		// Simple heuristic: skip binary files by checking extension or size (optional)
		if !isTextFile(path) {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return nil // ignore unreadable files
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNum := 0
		for scanner.Scan() {
			lineNum++
			line := scanner.Text()
			for _, re := range regexes {
				if re.MatchString(line) {
					exp, ok := internal.FileInspectionPatternExplanations[re.String()]
					if ok {
						results = append(results, fmt.Sprintf("%s:%d: matches '%s' - %s", path, lineNum, re.String(), exp))
					} else {
						// If no explanation is found, just report the match
						results = append(results, fmt.Sprintf("%s:%d: matches '%s'", path, lineNum, re.String()))
					}
				}
			}
		}

		return nil
	})

	return results, err
}

// isTextFile is a simple helper to guess if file is text (based on extension)
func isTextFile(filename string) bool {
	lower := strings.ToLower(filename)
	textExtensions := []string{
		".go", ".js", ".ts", ".jsx", ".tsx", ".php", ".html", ".css", ".scss", ".json", ".yaml", ".yml", ".md", ".txt",
	}

	for _, ext := range textExtensions {
		if strings.HasSuffix(lower, ext) {
			return true
		}
	}
	return false
}
