# AGENTS.md — Repository Instructions for AI Coding Agents

## Project Purpose
A lightweight CLI expense tracker written in Go (`go 1.27.3`) that stores expense records in a local JSON file, calculates totals, and exports reports in `text` and `csv` formats.

## Architecture & File Map
- `main.go`: CLI entry point; parses `os.Args` and delegates actions to `Store`.
- `tracker.go`: Core domain model (`Expense`), JSON persistence (`Store`), and report formatting (`Exporter` implementations).
- `tracker_test.go`: Automated unit tests using isolated temporary directories (`t.TempDir()`).
- `go.mod`: Module definition (`expense-tracker`, Go `1.27.3`), standard library only.

## Setup & Safe Commands
Agents are permitted to run the following read/verify commands freely:
- Check Go environment: `go version`
- Static analysis: `go vet ./...`
- Run test suite: `go test -v ./...`
- Build binary locally: `go build -o expense-tracker .`

## Coding Conventions & Test Expectations
- Keep dependencies strictly limited to the Go standard library unless explicitly approved.
- Format all Go files using `gofmt`.
- Never write test data to `expenses.json` in the repository root; always use `t.TempDir()` in tests.
- Every structural or behavioral change must keep all existing tests passing and include unit tests for new logic.

## Forbidden Actions
- **No Secrets:** Never hardcode API keys, credentials, tokens, or real personal data.
- **No Production Deployments:** Do not configure or execute deployment pipelines or external network calls.
- **No Destructive Commands:** Never run `rm -rf`, `git push --force`, `git reset --hard`, or delete `expenses.json`.
- **No Unreviewed Dependencies:** Do not run `go get` or modify `go.mod` to add external packages without explicit human approval.
