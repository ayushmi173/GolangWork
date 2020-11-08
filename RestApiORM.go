package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type Booking struct {
	Id      int    `json:"id"`
	User    string `json:"user"`
	Members int    `json:"members"`
}

func handleRequests() {
	log.Println("Starting development server at http://127.0.0.1:3000/")
	log.Println("Quit the server with CONTROL-C.")

	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)

	myRouter.HandleFunc("/booking/{id}", returnSingleBooking).Methods("GET")

	myRouter.HandleFunc("/new", createNewBooking).Methods("POST")

	myRouter.HandleFunc("/all", returnAllBookings).Methods("GET")

	myRouter.HandleFunc("/update/{id}", updateBooking).Methods("PUT")

	myRouter.HandleFunc("/delete/{id}", deleteBooking).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", myRouter))
}

func main() {
	// Please define your user name and password for my sql.
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True")
	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function

	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}

	db.AutoMigrate(&Booking{})
	handleRequests()
	defer db.Close()
}

func homePage(w http.ResponseWriter, r *http.Request) 
{
	fmt.Fprintf(w, "Welcome to Booking Home Page!")
	fmt.Println("Endpoint Hit: Booking Home Page")
}

func createNewBooking(w http.ResponseWriter, r *http.Request) {

	//return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var booking Booking
	json.Unmarshal(reqBody, &booking)
	db.Create(&booking)
	json.NewEncoder(w).Encode(booking)
}

func returnAllBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookings := []Booking{}
	//retrieve all bookings from db
	db.Find(&bookings)
	json.NewEncoder(w).Encode(bookings)
}

func returnSingleBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	key := mux.Vars(r)["id"]
	bookings := []Booking{}
	db.Find(&bookings)
	for _, booking := range bookings {
		//string to int
		s, err := strconv.Atoi(key)
		if err == nil {
			if booking.Id == s {
				fmt.Println(booking)
				json.NewEncoder(w).Encode(booking)
			}
		}
	}
}

func updateBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	bookings := []Booking{}
	db.Find(&bookings)

	//string to int
	s, err := strconv.Atoi(key)
	length := 0

	for _, booking := range bookings {
		if err == nil {
			if booking.Id == s {
				reqBody, _ := ioutil.ReadAll(r.Body)
				json.Unmarshal(reqBody, &booking)
				fmt.Println(booking)
				db.Save(booking)
				json.NewEncoder(w).Encode(booking)
				length = 1
			}
		}
	}
	if length == 0 {
		json.NewEncoder(w).Encode("Id doesn't exist")
		return
	}
}

func deleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	s, err := strconv.Atoi(key)
	if err == nil {
		db.Table("bookings").Where("id= ?", s).Delete(&Booking{})
		fmt.Fprintln(w, "Delete data of id ", s)
	}
}
