package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type User struct {
	ID int
	Email string
	Active bool
	DOB time.Time
	Score float64
}

func main() {

	f, err := os.Open("testdata/users.csv")
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(f)

	// skip headers
	_, _ = r.Read()

	var users []User

	for {
		record, err := r.Read()
		if err != nil {
			break
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Errorf("could not read id, err=", err)
		}
		active, err := strconv.ParseBool(record[2])
		if err != nil {
			fmt.Errorf("could not read active, err=", err)
		}
		dob, err := time.Parse("2006-01-02", record[3])
		if err != nil {
			fmt.Errorf("could not read birth_date, err=", err)
		}
		score, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			fmt.Errorf("could not read score, err=", err)
		}

		u := User{
			ID: id,
			Email: record[1],
			Active: active,
			DOB: dob,
			Score: score,
		}

		users = append(users, u)
	}

	fmt.Println("CSV parsing complete.")
	fmt.Printf("Read %d records\n\n", len(users))

	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}
}
