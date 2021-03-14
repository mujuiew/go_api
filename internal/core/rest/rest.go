package rest

import (
	"encoding/json"
	"net/http"
	"time"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var in Input
	_ = json.NewDecoder(r.Body).Decode(&in)

	date := in.CalDate
	t, _ := time.Parse("2006-01-02", date)
	caldate := t.Format("2006-01-02")
	proname, err := FineProname(caldate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	output := Output{proname, in.AccountNumber}
	js, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}
