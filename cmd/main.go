package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func readyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ready")
	fmt.Fprintf(w, "Api is ready")
}

func getEntryHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get all entries")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Register")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("login")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", readyHandler)
	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/register", registerHandler)
	log.Fatal(http.ListenAndServe(":8000", router))
}
