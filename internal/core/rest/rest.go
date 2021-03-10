package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Handle() {
	InitDB()
	r := mux.NewRouter()
	r.HandleFunc("/test/api", homePage).Methods("POST")
	http.ListenAndServe(":8010", r)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var in Input
	_ = json.NewDecoder(r.Body).Decode(&in)

	date := in.CalDate
	t, _ := time.Parse("2006-01-02", date)
	caldate := t.Format("2006-01-02")
	proname := FineProname(caldate)
	fmt.Fprintln(w, proname)
	// output := Output{proname, in.AccountNumber}
	// js, err := json.Marshal(output)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(js)
}
