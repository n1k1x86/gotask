package db

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/lib/pq"
)

func Init(host string, user string, password string, dbname string, sslmode string) *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		host, user, password, dbname, sslmode))

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS taskgroups(id SERIAL PRIMARY KEY, name VARCHAR(255))")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS tasks(id SERIAL PRIMARY KEY, title VARCHAR(255) NOT NULL, description TEXT NOT NULL, groupId INTEGER REFERENCES taskgroups (id), FOREIGN KEY (groupId) REFERENCES taskgroups (id))")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
