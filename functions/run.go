package functions

import (
	"log"
	"net/http"
)

// Run starts the service.
func Run() {

	http.HandleFunc("/register", Register)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/contacted", Contacted)
	http.HandleFunc("/query", Query)
	http.HandleFunc("/dump", Dump)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
