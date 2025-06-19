package internal

import (
	"regexp"
	"strings"
)

var CommitMessagePatternExplanations = map[string]string{
	`(?i)^fix$`: "Avoid using 'fix' alone as a commit message. Use descriptive messages that explain the fix.",
	`(?i)^wip`:  "Work-in-progress commits should be avoided in shared branches to keep history clean.",
	`(?i)final`: "Avoid vague 'final' commit messages. Be specific about what was finalized.",
	`(?i)temp`:  "Temporary commits should be squashed or removed before merging.",
	`(?i)debug`: "Debug-related commits should not be part of the main history.",
}

var FileInspectionPatternExplanations = map[string]string{
	`console\.log`:         "Leaving `console.log` statements in production code can leak information and clutter logs.",
	`debugger`:             "The `debugger` statement pauses JS execution; it should be removed before shipping.",
	`dd`:                   "`dd` (dump and die) is a debugging helper that should not be left in committed code.",
	`dump`:                 "Debug dumps may expose sensitive information or clutter output.",
	`TODO`:                 "TODO comments should be addressed or tracked outside code to avoid technical debt.",
	`print_r\(`:            "PHP print_r statements should not be left in committed code.",
	`var_dump\(`:           "PHP var_dump statements should be removed before committing.",
	`System\.out\.println`: "Java System.out.println calls should be removed before committing.",
	`logger\.debug`:        "Debug-level logging statements should be avoided in production commits.",
	`UNUSED`:               "Unused code or variables should be cleaned up before committing.",
	`DEAD CODE`:            "Dead code should be removed to keep the codebase clean.",
	`if\s*\(\s*false\s*\)`: "Conditional blocks that never execute should be removed.",
	`#if\s+0`:              "Disabled code blocks using preprocessor directives should be cleaned up.",
	`\b(xit|it\.skip|describe\.skip|test\.only|focus:)\b`: "Focused or skipped tests should not be committed.",
}

func DocLinkForPattern(pattern, patternType string) string {
	baseURL := "https://rickardahlstedt.github.io/gitgeist/"
	subdir := map[string]string{
		"commit": "cmt-msg",
		"file":   "file",
	}[patternType]

	// Remove (?i) and similar flags
	pattern = regexp.MustCompile(`(?i)\(\?[a-z]+\)`).ReplaceAllString(pattern, "")

	// Extract last alphanumeric segment
	re := regexp.MustCompile(`[a-zA-Z0-9]+`)
	words := re.FindAllString(pattern, -1)
	if len(words) == 0 {
		return baseURL + "index.html"
	}
	filename := strings.ToLower(words[len(words)-1]) + ".html"

	return baseURL + subdir + "/" + filename
}
