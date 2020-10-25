package main

//https://github.com/bradtraversy/go_restapi/blob/master/main.go
import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//This is global struct(Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init books as a slice
var books []Book

//Main Function
func main() {

	//init router
	r := mux.NewRouter()

	//add hardcoded data
	books = append(books, Book{ID: "1", Isbn: "2743215", Title: "The Great war", Author: &Author{Firstname: "Daryl", Lastname: "lasol's"}})
	books = append(books, Book{ID: "2", Isbn: "97654334", Title: "Rempol Sketch", Author: &Author{Firstname: "Jubian", Lastname: "Rescos"}})

	//Route handler for the end points
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/book", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/book", updateBook).Methods("PUT")

	//Start the server
	log.Fatal(http.ListenAndServe(":3000", r))
}

//Get all books from books slice
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//get single book from the books slice
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get Params

	//loop through the books and find id from the params
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//Add a new Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var b Book //taking struct in book variable
	json.NewDecoder(r.Body).Decode(&b)
	b.ID = strconv.Itoa(rand.Intn(10000000)) // temp id
	books = append(books, b)
	json.NewEncoder(w).Encode(b)
}

//update book

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			break
		}
	}
	json.NewEncoder(w).Encode(book)
}
