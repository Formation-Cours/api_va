package main

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type User struct {
	ID        uint   `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
}

func (u *User) insert(db *sql.DB) error {
	query := "INSERT INTO user(firstname, lastname, email) VALUES (?, ?, ?)"
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}(stmt)
	res, err := stmt.ExecContext(ctx, u.Firstname, u.Lastname, u.Email)
	if err != nil {
		log.Printf("Error %s when inserting row into user table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d user created ", rows)
	return nil
}
func (u *User) findAll(db *sql.DB) ([]User, error) {
	query := "SELECT * FROM user"
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return []User{}, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email); err != nil {
			return []User{}, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return []User{}, err
	}
	return users, nil
}
