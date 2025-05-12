# GShell - A Minimalist Shell in Go

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)

GShell is a lightweight command-line shell implemented in Go, designed for educational purposes and personal experimentation with shell internals.

## Features

### Current Features
✅ Basic command execution (e.g., `ls`, `pwd`, `echo`)
✅ Change directory (`cd`) command
✅ Command pipelines (`|`)
✅ Colored output and error messages
✅ Custom shell prompt with current directory

### Planned Features
🔜 Command history navigation
🔜 Background process execution (`&`)
🔜 Chained commands (`;`, `&&`)
🔜 Custom aliases
🔜 Shell configuration via `.gshellrc`
🔜 Built-in commands (`help`, `clear`, custom `ls`/`cat`)

## Installation

1. Ensure you have Go 1.20+ installed
2. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/gshell.git
   cd gshell
   ```
3. Build and run:
   ```bash
   go build -o gshell .
   ./gshell
   ```

## Usage

```
🚀 /current/directory > ls -la
🚀 /current/directory > cd ~/projects
🚀 ~/projects > echo "Hello" | grep "H"
```

## Project Structure

```
gshell/
├── command/           # Command execution logic
│   ├── exec.go        # Main command execution
│   └── cd.go          # Directory changing
├── prompt/            # Shell interface
│   ├── colors.go      # Color handling
│   └── prompt.go      # Prompt display
├── main.go            # Entry point
└── README.md          # This file
```

---

**Note**: This is an educational project. Not recommended for production use.
