package main

import (
	"go-user-service/db"
	"go-user-service/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting user service...")
	db.Connect()

	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.SaveUser).Methods("POST")
	r.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	r.HandleFunc("/users/{email}", handlers.UpdateUserByEmail).Methods("PUT")
	r.HandleFunc("/users/{email}", handlers.DeleteUserByEmail).Methods("DELETE")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
