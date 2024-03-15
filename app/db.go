package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		username VARCHAR(64) NOT NULL UNIQUE,
		password VARCHAR(64) NOT NULL,
		register_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY,
		name VARCHAR(64) NOT NULL UNIQUE,
		description VARCHAR(64) NOT NULL,
		author_id INTEGER,
		is_private BOOLEAN DEFAULT 0,
		publicate_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (author_id) REFERENCES users(id)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func register(db *sql.DB, username string, password1 string, password2 string) bool {
	if password1 != password2 {
		log.Println(red("Passwords do not match"))
		return false
	}

	hashedPassword := getHash(password1)

	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func login(db *sql.DB, username string, password string) bool {
	var storedHash string

	err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(red("User not founded"))
		} else {
			log.Println(red(err))
		}
		return false
	}
	if compareHash(password, storedHash) {
		return true
	} else {
		log.Println(red("Wrong password"))
		return false
	}
}

func getAuthorIdByName(db *sql.DB, author string) int {
	var author_id int

	err := db.QueryRow("SELECT id FROM users WHERE username = ?", author).Scan(&author_id)
	if err != nil {
		log.Println(red(err))
		return 0
	}
	return author_id
}

func addNote(db *sql.DB, author string, noteName string, description string, isPrivate bool) bool {
	author_id := getAuthorIdByName(db, author)

	_, err := db.Exec("INSERT INTO notes (name, description, author_id, is_private) VALUES (?, ?, ?, ?)", noteName, description, author_id, isPrivate)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func deleteNote(db *sql.DB, author string, noteName string) bool {
	author_id := getAuthorIdByName(db, author)

	_, err := db.Exec("DELETE FROM notes WHERE name = ? AND author_id = ?", noteName, author_id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func getNotes(db *sql.DB, author string) map[string]map[string]interface{} {
	author_id := getAuthorIdByName(db, author)
	notes := make(map[string]map[string]interface{})

	rows, err := db.Query("SELECT name, description FROM notes WHERE author_id = ?", author_id)
	if err != nil {
		log.Println(err)
		return nil
	}

	for rows.Next() {
		var name string
		var description string

		err = rows.Scan(&name, &description)
		if err != nil {
			log.Println(err)
			break
		}

		notes[name] = map[string]interface{}{
			"author":      author,
			"description": description,
		}
	}
	rows.Close()

	rows, err = db.Query("SELECT name, description FROM notes WHERE is_private = false")
	if err != nil {
		log.Println(err)
		return nil
	}

	for rows.Next() {
		var name string
		var description string

		err = rows.Scan(&name, &description)
		if err != nil {
			log.Println(err)
			break
		}

		notes[name] = map[string]interface{}{
			"author":      author,
			"description": description,
		}
	}
	rows.Close()

	return notes
}
