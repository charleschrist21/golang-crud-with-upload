package main

import (
	"golang-rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Init Router
	r := mux.NewRouter()

	// arrange our route

	//Books
	r.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods("DELETE")

	//upload
	r.HandleFunc("/upload", controllers.UploadFile).Methods("POST")

	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))
}
