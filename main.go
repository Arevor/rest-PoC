package main

// make sure DB is running, i.e
// systemctl start postgresql.service

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	tmpDB, err := sql.Open("postgres", "user=postgres dbname=kanban_desu sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/cards", handleInsertCard).Methods("POST")
	r.HandleFunc("/cards", handleUpdateCard).Methods("PUT")
	r.HandleFunc("/cards", handleGetCards).Methods("GET")
	r.HandleFunc("/card/{id:[0-9]+}", handleDeleteCard).Methods("DELETE")
	r.HandleFunc("/card/{id:[0-9]+}", handleGetCard).Methods("GET")

	r.HandleFunc("/card/{id:[0-9]+}/comment", handlePostComment).Methods("POST")
	r.HandleFunc("/card/{id:[0-9]+}/comment/{id2:[0-9]+}", handlePutComment).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}
