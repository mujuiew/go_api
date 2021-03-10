package rest

import (
	"fmt"
	"log"
	"net/http"
)

func Handle() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
