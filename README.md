# Expense Tracker CLI

A lightweight CLI expense tracker built with Go 1.27.3 using only the Go standard library.

## Features
- Add expenses with input validation (`add <title> <category> <amount>`)
- Calculate total expenses (`total`)
- Export reports via the Strategy pattern (`export <text|csv>`)

## Setup & Usage
    go run . add "Coffee" "Food" 850
    go run . total
    go run . export text
    go run . export csv

## Verification & Build
    go vet ./...
    go test -v ./...
    go build -o expense-tracker .
