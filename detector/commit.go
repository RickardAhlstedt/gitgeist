package detector

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/RickardAhlstedt/gitgeist/internal"
)

func GetLastCommitMessage() (string, error) {
	cmd := exec.Command("git", "log", "-1", "--pretty=%B")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to get last commit message: %w", err)
	}
	return strings.TrimSpace(out.String()), nil
}

func AnalyzeCommitMessage(message string, patterns []string) []string {
	var warnings []string

	if len(message) < 10 {
		warnings = append(warnings, "Commit message too short")
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		if re.MatchString(message) {
			exp, ok := internal.CommitMessagePatternExplanations[pattern]
			if ok {
				warnings = append(warnings, fmt.Sprintf("Pattern '%s' found: %s (%s)", pattern, exp, internal.DocLinkForPattern(re.String(), "commit")))
			} else {
				warnings = append(warnings, fmt.Sprintf("Pattern '%s' found", pattern))
			}
		}
	}
	return warnings
}
