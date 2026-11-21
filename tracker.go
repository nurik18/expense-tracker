package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Expense struct {
	Title    string  `json:"title"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
}

// Exporter defines the Strategy interface for formatting expense reports.
type Exporter interface {
	Format(items []Expense, total float64) string
}

type TextExporter struct{}

func (e TextExporter) Format(items []Expense, total float64) string {
	var sb strings.Builder
	for _, item := range items {
		sb.WriteString(fmt.Sprintf("- %s [%s]: %.2f\n", item.Title, item.Category, item.Amount))
	}
	sb.WriteString(fmt.Sprintf("Total: %.2f\n", total))
	return sb.String()
}

type CSVExporter struct{}

func (e CSVExporter) Format(items []Expense, total float64) string {
	var sb strings.Builder
	sb.WriteString("title,category,amount\n")
	for _, item := range items {
		sb.WriteString(fmt.Sprintf("%s,%s,%.2f\n", item.Title, item.Category, item.Amount))
	}
	return sb.String()
}

var exporters = map[string]Exporter{
	"text": TextExporter{},
	"csv":  CSVExporter{},
}

type Store struct {
	FilePath string
	Items    []Expense
}

func NewStore(path string) (*Store, error) {
	s := &Store{FilePath: path}
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return s, nil
		}
		return nil, fmt.Errorf("read store file: %w", err)
	}
	if len(strings.TrimSpace(string(data))) == 0 {
		return s, nil
	}
	if err := json.Unmarshal(data, &s.Items); err != nil {
		return nil, fmt.Errorf("parse store json: %w", err)
	}
	return s, nil
}

func (s *Store) Add(title, category string, amount float64) error {
	if strings.TrimSpace(title) == "" {
		return errors.New("title cannot be empty")
	}
	if strings.TrimSpace(category) == "" {
		return errors.New("category cannot be empty")
	}
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	s.Items = append(s.Items, Expense{
		Title:    strings.TrimSpace(title),
		Category: strings.TrimSpace(category),
		Amount:   amount,
	})

	data, err := json.MarshalIndent(s.Items, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal expenses: %w", err)
	}
	if err := os.WriteFile(s.FilePath, data, 0644); err != nil {
		return fmt.Errorf("write store file: %w", err)
	}
	return nil
}

func (s *Store) Total() float64 {
	var sum float64
	for _, item := range s.Items {
		sum += item.Amount
	}
	return sum
}

func (s *Store) Export(format string) (string, error) {
	exporter, ok := exporters[format]
	if !ok {
		return "", fmt.Errorf("unsupported format: %s", format)
	}
	return exporter.Format(s.Items, s.Total()), nil
}
