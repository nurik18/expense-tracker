package main

import (
	"path/filepath"
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
