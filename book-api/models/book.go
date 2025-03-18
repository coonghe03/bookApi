package models

import "encoding/json"

// Book represents a book entity
type Book struct {
    BookID          string  `json:"bookId"`
    AuthorID        string  `json:"authorId"`
    PublisherID     string  `json:"publisherId"`
    Title           string  `json:"title"`
    PublicationDate string  `json:"publicationDate"`
    ISBN            string  `json:"isbn"`
    Pages           int     `json:"pages"`
    Genre           string  `json:"genre"`
    Description     string  `json:"description"`
    Price           float64 `json:"price"`
    Quantity        int     `json:"quantity"`
}

// Convert a Book struct to JSON
func (b *Book) ToJSON() string {
    jsonBytes, _ := json.MarshalIndent(b, "", "  ")
    return string(jsonBytes)
}
