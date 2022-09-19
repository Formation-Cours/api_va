package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	db = connecDB()
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			var user User
			users, err := user.findAll(db)
			if err != nil {
				log.Println(err)
				return
			}
			err = json.NewEncoder(w).Encode(users)
			if err != nil {
				log.Println(err)
				return
			}
			return
		}
	})

	log.Fatalln(http.ListenAndServe("localhost:8080", mux))
}
