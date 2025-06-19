package internal

var CommitMessagePatternExplanations = map[string]string{
	`(?i)^fix$`: "Avoid using 'fix' alone as a commit message. Use descriptive messages that explain the fix.",
	`(?i)^wip`:  "Work-in-progress commits should be avoided in shared branches to keep history clean.",
	`(?i)final`: "Avoid vague 'final' commit messages. Be specific about what was finalized.",
	`(?i)temp`:  "Temporary commits should be squashed or removed before merging.",
	`(?i)debug`: "Debug-related commits should not be part of the main history.",
}

var FileInspectionPatternExplanations = map[string]string{
	`console\.log`: "Leaving `console.log` statements in production code can leak information and clutter logs.",
	`debugger`:     "The `debugger` statement pauses JS execution; it should be removed before shipping.",
	`dd`:           "`dd` (dump and die) is a debugging helper that should not be left in committed code.",
	`dump`:         "Debug dumps may expose sensitive information or clutter output.",
	`TODO`:         "TODO comments should be addressed or tracked outside code to avoid technical debt.",
}
