package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "book-api/models"
)

// Get all books
func getBooksHandler(w http.ResponseWriter, r *http.Request) {
    books, err := models.ReadBooks()
    if err != nil {
        http.Error(w, "Error reading books", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

// Get a book by ID
func getBookByIDHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookID := vars["id"]

    books, err := models.ReadBooks()
    if err != nil {
        http.Error(w, "Error reading books", http.StatusInternalServerError)
        return
    }

    for _, book := range books {
        if book.BookID == bookID {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(book)
            return
        }
    }

    http.Error(w, "Book not found", http.StatusNotFound)
}

// Add a new book
func addBookHandler(w http.ResponseWriter, r *http.Request) {
    var newBook models.Book
    err := json.NewDecoder(r.Body).Decode(&newBook)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    books, err := models.ReadBooks()
    if err != nil {
        http.Error(w, "Error reading books", http.StatusInternalServerError)
        return
    }

    books = append(books, newBook)

    err = models.WriteBooks(books)
    if err != nil {
        http.Error(w, "Error saving book", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newBook)
}

// Update a book by ID
func updateBookHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookID := vars["id"]

    var updatedBook models.Book
    err := json.NewDecoder(r.Body).Decode(&updatedBook)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    books, err := models.ReadBooks()
    if err != nil {
        http.Error(w, "Error reading books", http.StatusInternalServerError)
        return
    }

    for i, book := range books {
        if book.BookID == bookID {
            books[i] = updatedBook
            err = models.WriteBooks(books)
            if err != nil {
                http.Error(w, "Error saving book", http.StatusInternalServerError)
                return
            }
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(updatedBook)
            return
        }
    }

    http.Error(w, "Book not found", http.StatusNotFound)
}

// Delete a book by ID
func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookID := vars["id"]

    books, err := models.ReadBooks()
    if err != nil {
        http.Error(w, "Error reading books", http.StatusInternalServerError)
        return
    }

    for i, book := range books {
        if book.BookID == bookID {
            books = append(books[:i], books[i+1:]...)
            err = models.WriteBooks(books)
            if err != nil {
                http.Error(w, "Error saving books", http.StatusInternalServerError)
                return
            }
            w.WriteHeader(http.StatusOK)
            fmt.Fprintf(w, "Book with ID %s deleted", bookID)
            return
        }
    }

    http.Error(w, "Book not found", http.StatusNotFound)
}

// Main function
func main() {
    router := mux.NewRouter()

    router.HandleFunc("/books", getBooksHandler).Methods("GET")
    router.HandleFunc("/books/{id}", getBookByIDHandler).Methods("GET")
    router.HandleFunc("/books", addBookHandler).Methods("POST")
    router.HandleFunc("/books/{id}", updateBookHandler).Methods("PUT")
    router.HandleFunc("/books/{id}", deleteBookHandler).Methods("DELETE")

    fmt.Println("Server is running on port 9090...")
    http.ListenAndServe(":9090", router)
}
