package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestAddAndTotal(t *testing.T) {
	tempFile := filepath.Join(t.TempDir(), "test_expenses.json")
	store := NewStore(tempFile)

	store.Add("Lunch", "Food", 1500)
	store.Add("Bus", "Transport", 200)

	if got := store.Total(); got != 1700 {
		t.Errorf("expected total 1700, got %.2f", got)
	}
}

func TestExportFormats(t *testing.T) {
	tempFile := filepath.Join(t.TempDir(), "test_export.json")
	store := NewStore(tempFile)
	store.Add("Coffee", "Food", 850)

	textOut, err := store.Export("text")
	if err != nil || !strings.Contains(textOut, "Total: 850.00") {
		t.Fatalf("unexpected text export: %v, %s", err, textOut)
	}

	csvOut, err := store.Export("csv")
	if err != nil || !strings.Contains(csvOut, "Coffee,Food,850.00") {
		t.Fatalf("unexpected csv export: %v, %s", err, csvOut)
	}
}
