package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Expense struct {
	Title    string  `json:"title"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
}

type Store struct {
	FilePath string
	Items    []Expense
}

func NewStore(path string) *Store {
	s := &Store{FilePath: path}
	data, err := os.ReadFile(path)
	if err == nil {
		_ = json.Unmarshal(data, &s.Items)
	}
	return s
}

func (s *Store) Add(title, category string, amount float64) {
	s.Items = append(s.Items, Expense{
		Title:    title,
		Category: category,
		Amount:   amount,
	})
	data, _ := json.MarshalIndent(s.Items, "", "  ")
	_ = os.WriteFile(s.FilePath, data, 0644)
}

func (s *Store) Total() float64 {
	var sum float64
	for _, item := range s.Items {
		sum += item.Amount
	}
	return sum
}

func (s *Store) Export(format string) (string, error) {
	if format == "text" {
		var sb strings.Builder
		for _, item := range s.Items {
			sb.WriteString(fmt.Sprintf("- %s [%s]: %.2f\n", item.Title, item.Category, item.Amount))
		}
		sb.WriteString(fmt.Sprintf("Total: %.2f\n", s.Total()))
		return sb.String(), nil
	} else if format == "csv" {
		var sb strings.Builder
		sb.WriteString("title,category,amount\n")
		for _, item := range s.Items {
			sb.WriteString(fmt.Sprintf("%s,%s,%.2f\n", item.Title, item.Category, item.Amount))
		}
		return sb.String(), nil
	}
	return "", fmt.Errorf("unsupported format: %s", format)
}
