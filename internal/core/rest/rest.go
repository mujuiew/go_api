package rest

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Handle() {
	InitDB()
	// http.HandleFunc("/", homePage)
	// log.Fatal(http.ListenAndServe(":8001", nil))

	r := mux.NewRouter()
	r.HandleFunc("/test/api", homePage).Methods("POST")
	http.ListenAndServe(":8080", r)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	var in Input

	date := in.CalDate
	t, _ := time.Parse("2006-01-02", date)
	caldate := t.Format("2006-01-02")
	proname := FineProname(caldate)
	inrate := Fineinrate(proname)
	pmt := InsertAc(in.AccountNumber, inrate, in.DisbursementAmount, in.NumberOfPayment)

	inrate = math.Round(inrate*100) / 100
	pmt = math.Round(pmt*100) / 100

	output := Output{proname, inrate, in.AccountNumber, pmt}
	js, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Hello")

	w.Write(js)
}
