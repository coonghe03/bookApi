package models

import (
	"testing"
)

func TestSearchBooksConcurrent(t *testing.T) {
    // Sample test data
    mockBooks := []Book{
        {
            BookID:      "1",
            Title:       "Go Programming",
            Description: "A book about Go programming",
        },
        {
            BookID:      "2",
            Title:       "Python Programming",
            Description: "A book about Python",
        },
    }

    // Mock read function
    mockRead := func() ([]Book, error) {
        return mockBooks, nil
    }

    results, err := SearchBooksConcurrent("go", mockRead)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if len(results) != 1 {
        t.Errorf("Expected 1 result, got %d", len(results))
    }

    if results[0].BookID != "1" {
        t.Errorf("Expected BookID '1', got '%s'", results[0].BookID)
    }
}
