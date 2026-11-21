package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAddAndTotal(t *testing.T) {
	tempFile := filepath.Join(t.TempDir(), "test_expenses.json")
	store, err := NewStore(tempFile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := store.Add("Lunch", "Food", 1500); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := store.Add("Bus", "Transport", 200); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := store.Total(); got != 1700 {
		t.Errorf("expected total 1700, got %.2f", got)
	}
}

func TestExportFormats(t *testing.T) {
	tempFile := filepath.Join(t.TempDir(), "test_export.json")
	store, err := NewStore(tempFile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := store.Add("Coffee", "Food", 850); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	textOut, err := store.Export("text")
	if err != nil || !strings.Contains(textOut, "Total: 850.00") {
		t.Fatalf("unexpected text export: %v, %s", err, textOut)
	}

	csvOut, err := store.Export("csv")
	if err != nil || !strings.Contains(csvOut, "Coffee,Food,850.00") {
		t.Fatalf("unexpected csv export: %v, %s", err, csvOut)
	}

	if _, err := store.Export("xml"); err == nil {
		t.Errorf("expected error for unsupported format, got nil")
	}
}

func TestAddValidation(t *testing.T) {
	tempFile := filepath.Join(t.TempDir(), "test_validation.json")
	store, err := NewStore(tempFile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := store.Add("", "Food", 500); err == nil {
		t.Errorf("expected error for empty title, got nil")
	}
	if err := store.Add("Taxi", "", 500); err == nil {
		t.Errorf("expected error for empty category, got nil")
	}
	if err := store.Add("Taxi", "Transport", -100); err == nil {
		t.Errorf("expected error for negative amount, got nil")
	}
	if err := store.Add("Taxi", "Transport", 0); err == nil {
		t.Errorf("expected error for zero amount, got nil")
	}
}

func TestCorruptedJSONHandling(t *testing.T) {
	tempFile := filepath.Join(t.TempDir(), "corrupted.json")
	if err := os.WriteFile(tempFile, []byte("{invalid-json"), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	if _, err := NewStore(tempFile); err == nil {
		t.Errorf("expected error when loading corrupted JSON, got nil")
	}
}
