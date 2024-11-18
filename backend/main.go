package main

import (
	_ "database/sql"
	"log"
	"net/http"
)

func createTransaction(w http.ResponseWriter, r *http.Request) {

}

func setupDB() {

}

func main() {
	setupDB()

	router := http.NewServeMux()
	router.HandleFunc("/transactions", createTransaction)

	err := http.ListenAndServe(":3333", router)
	if err != nil {
		log.Fatal(err)
	}

}
