package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

//Book Struct (Model)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init Books var as a slice Book Struct
var books []Book

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

//Get A Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get Params

	//Loop through books and find book with id
	for _, item := range books {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(&Book{})
	if err != nil {
		return
	}
}

//Create A Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe

	books = append(books, book)
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		return
	}
}

//Update A Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"] // Mock ID - not safe

			books = append(books, book)
			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

//Delete A Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

func main() {
	//Init Router
	r := mux.NewRouter()

	//Mock Data
	books = append(books, Book{
		ID:    "1",
		Isbn:  "448112",
		Title: "Book One",
		Author: &Author{
			Firstname: "Asil",
			Lastname:  "Hazbik",
		},
	})
	books = append(books, Book{
		ID:    "2",
		Isbn:  "448116",
		Title: "Book Two",
		Author: &Author{
			Firstname: "Theo",
			Lastname:  "Vasilis",
		},
	})

	//Route Handlers / End Points
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
