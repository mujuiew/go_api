package rest

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5503
	user     = "postgres"
	password = "postgres"
	dbname   = "payment"
)

var db *sql.DB

// var pro Promotion
// var rate Rate

func InitDB() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Println("\n============== Error db.Ping ====================")
		log.Panic(err.Error())
	}

	fmt.Println("Successfully connected! DB")
}

// FineProname ...
func FineProname(date string) (string, error) {
	var pro Promotion
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, _ = sql.Open("postgres", psqlInfo)

	rows, err := db.Query(`SELECT promotion_name FROM "promotion" WHERE '` + date + `'  between start_date and end_date `)
	if err != nil {
		return "fail Query ", err
	}
	for rows.Next() {
		if err := rows.Scan(&pro.PromotionName); err != nil {
			log.Fatal(err)
		}
	}
	return pro.PromotionName, nil
}

// // Fineinrate ...
// func Fineinrate(pro string) float64 {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	db, _ := sql.Open("postgres", psqlInfo)

// 	rows, _ := db.Query(`SELECT interest_rate FROM "Rate" WHERE promotion_name = '` + pro + `'`)
// 	for rows.Next() {
// 		if err := rows.Scan(&rate.InterestRate); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	inrate := rate.InterestRate
// 	return inrate
// }

// // InsertAc ...
// func InsertAc(acc float64, i float64, dis float64, num float64) float64 {
// 	i = i / 100 / 12
// 	pmt := dis / ((1 - (1 / math.Pow((1+i), num))) / i)

// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	err = db.Ping()
// 	queryStr := `INSERT INTO public."Account"(
// 		account_number, installment_amount)
// 		VALUES ($1, $2)`
// 	_, err = db.Exec(queryStr, acc, pmt)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return pmt
// }
