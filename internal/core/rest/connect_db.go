package rest

import (
	"database/sql"
	"fmt"
	"log"
	"math"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "payment"
)

var db *sql.DB

var pro Promotion
var rate Rate

func InitDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

// FineProname ...
func FineProname(date string) string {
	rows, _ := db.Query(`SELECT promotion_name FROM "Promotion" WHERE '` + date + `'  between start_date and end_date `)
	for rows.Next() {
		if err := rows.Scan(pro.PromotionName); err != nil {
			log.Fatal(err)
		}
	}
	proname := pro.PromotionName

	return proname
}

// Fineinrate ...
func Fineinrate(pro string) float64 {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, _ := sql.Open("postgres", psqlInfo)

	rows, _ := db.Query(`SELECT interest_rate FROM "Rate" WHERE promotion_name = '` + pro + `'`)
	for rows.Next() {
		if err := rows.Scan(&rate.InterestRate); err != nil {
			log.Fatal(err)
		}
	}
	inrate := rate.InterestRate
	return inrate
}

// InsertAc ...
func InsertAc(acc float64, i float64, dis float64, num float64) float64 {
	i = i / 100 / 12
	pmt := dis / ((1 - (1 / math.Pow((1+i), num))) / i)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	err = db.Ping()
	queryStr := `INSERT INTO public."Account"(
		account_number, installment_amount)
		VALUES ($1, $2)`
	_, err = db.Exec(queryStr, acc, pmt)
	if err != nil {
		panic(err)
	}
	return pmt
}
