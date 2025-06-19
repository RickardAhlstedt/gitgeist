# 👹 Gitgeist (a.k.a The Commit Wraith)

“Every time you force-push to main, a file vanishes… forever.”

---

## 📜 Lore

The Commit Wraith is an ancient, spectral entity bound to version control systems by a thousand cursed merge conflicts. It lurks in neglected branches, malformed commit messages, and rebased histories — waiting, watching.

It was born from the tears of a junior developer who accidentally deleted `production.env` on a Friday afternoon.

---

## 🧙 What is Gitgeist?

Gitgeist is a playful CLI tool written in Go that haunts your Git repositories when bad practices are detected. It scans your commits and code for sins such as:

- Suspicious commit messages (`fix`, `wip`, `final`, etc.)
- Debug statements left in code (`console.log`, `dd`, `dump`, `TODO`)

### TBD  

- Force-pushing to protected branches
- Other common git anti-patterns (configurable)

Gitgeist helps you enforce better habits by warning or blocking problematic commits.

---

## 🚀 Installation

1. Download or build the binary:

```bash
git clone https://github.com/RickardAhlstedt/gitgeist.git
cd gitgeist
go build -o gitgeist ./cmd
```

2. Move `gitgeist` to your `$PATH` (optional):

```bash
mv gitgeist /usr/local/bin/
```

3. Run `gitgeist` in your repository root to see warnings.

---

## ⚙️ Configuration

Gitgeist reads its config from:

```bash
~/.gitgeist/config.yaml
```

If it doesn't exist, Gitgeist will create a default config on first run, including patterns to check for in commit messages and files.

Example config:

```yaml
commit_message_patterns:
  - '(?i)^fix$'
  - '(?i)^wip'
  - '(?i)final'
  - '(?i)temp'
  - '(?i)debug'

file_inspection_patterns:
  - 'console\.log'
  - 'debugger'
  - 'dd'
  - 'dump'
  - 'TODO'
```

---

## 🪄 Usage

### Basic check

Run the tool in your Git repo directory:

```bash
gitgeist
```

It will:

- Analyze your last commit message for suspicious patterns
- Scan your codebase for debug or TODO statements
- Report warnings or confirm a clean state

### Install Git Hook

To install a pre-commit Git hook that runs Gitgeist automatically before every commit, run:

```bash
gitgeist install-hook
```

This will install a hook in `.git/hooks/pre-commit` that blocks commits if Gitgeist detects issues.

---

## 🎭 How It Works

- Scans your last commit message using configurable regex patterns.
- Recursively scans text files (like `.js`, `.go`, `.php`, `.ts`) for suspicious debug statements.
- Reports filename and line numbers of offending lines.
- Encourages good Git hygiene and cleaner commits.

---

## 📖 License

MIT License

---

Let me know if you want me to add usage examples, build instructions for different platforms, or more detailed configuration options!