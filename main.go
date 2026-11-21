package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	store := NewStore("expenses.json")

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [add <title> <category> <amount> | total | export <text|csv>]")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 5 {
			fmt.Println("Usage: add <title> <category> <amount>")
			return
		}
		amount, err := strconv.ParseFloat(os.Args[4], 64)
		if err != nil {
			fmt.Println("Invalid amount:", err)
			return
		}
		store.Add(os.Args[2], os.Args[3], amount)
		fmt.Println("Expense added.")
	case "total":
		fmt.Printf("Total expenses: %.2f\n", store.Total())
	case "export":
		format := "text"
		if len(os.Args) >= 3 {
			format = os.Args[2]
		}
		out, err := store.Export(format)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Print(out)
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
