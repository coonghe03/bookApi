package models

import (
	"encoding/json"
	"os"
	"strings"
	"sync"
)

// File path for storing books
const filePath = "data/books.json"

// ReadBooks reads books from the JSON file
func ReadBooks() ([]Book, error) {
    file, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    var books []Book
    json.Unmarshal(file, &books)
    return books, nil
}

// WriteBooks writes books to the JSON file
func WriteBooks(books []Book) error {
    bytes, err := json.MarshalIndent(books, "", "  ")
    if err != nil {
        return err
    }

    return os.WriteFile(filePath, bytes, 0644)
}

func SearchBooksConcurrent(keyword string, readBooksFunc func() ([]Book, error)) ([]Book, error) {
    books, err := readBooksFunc()
    if err != nil {
        return nil, err
    }

    keyword = strings.ToLower(keyword)
    var results []Book
    var wg sync.WaitGroup
    resultChan := make(chan Book, len(books))

    for _, book := range books {
        wg.Add(1)
        go func(b Book) {
            defer wg.Done()

            if strings.Contains(strings.ToLower(b.Title), keyword) ||
                strings.Contains(strings.ToLower(b.Description), keyword) {
                resultChan <- b
            }
        }(book)
    }

    go func() {
        wg.Wait()
        close(resultChan)
    }()

    for b := range resultChan {
        results = append(results, b)
    }

    return results, nil
}
