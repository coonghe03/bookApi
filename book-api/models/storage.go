package models

import (
    "encoding/json"
    "io/ioutil"
    "os"
)

// File path for storing books
const filePath = "data/books.json"

// ReadBooks reads books from the JSON file
func ReadBooks() ([]Book, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    bytes, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }

    var books []Book
    json.Unmarshal(bytes, &books)
    return books, nil
}

// WriteBooks writes books to the JSON file
func WriteBooks(books []Book) error {
    bytes, err := json.MarshalIndent(books, "", "  ")
    if err != nil {
        return err
    }

    return ioutil.WriteFile(filePath, bytes, 0644)
}

