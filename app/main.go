package main

import (
	"fmt"
	"log"
)

var username string

func banner() {
	fmt.Println(cyan("+-+-+-+-+ +-+-+-+-+-+-+-+\n") +
		green("|N|o|t|e| |S|e|r|v|i|c|e|\n") +
		cyan("+-+-+-+-+ +-+-+-+-+-+-+-+\n\n"))
}

func menu1() {
	fmt.Println(blue("=========================\n") +
		green("	Quotes Menu       \n") +
		blue("=========================\n") +
		"1. Login\n2. Register\n3. Exit\n" +
		blue("========================="))
}

func menu2() {
	fmt.Println(blue("=========================\n") +
		green("	Quotes Menu       \n") +
		blue("=========================\n") +
		"1. Get list of quotes\n2. Add note\n3. Delete note\n4. View this menu\n5. Exit\n" +
		blue("========================="))
}

func printNotes(notes map[string]map[string]interface{}) {
	if len(notes) == 0 {
		fmt.Println("No notes :(\n======")
		return
	}

	c := 1
	for name := range notes {
		fmt.Printf("%d\nAuthor: %s\n"+blue("=======\n")+"Note: %s\nDescription: %s\n"+blue("=======\n"), c, notes[name]["author"], name, notes[name]["description"])
		c++
	}
}

func main() {
	var choice int8
	var isAuthorized bool = false

	banner()

	db := createDB("./data.db")
	defer db.Close()

	for {
		menu1()
		fmt.Print("Choice: ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			fmt.Print("Username: ")
			username = input()

			fmt.Print("Password: ")
			password := input()

			status := login(db, username, password)

			if status {
				isAuthorized = true
				log.Println("Login successful")
			} else {
				log.Println("Login failed")
			}
		case 2:
			fmt.Print("Username: ")
			username = input()

			fmt.Print("Password: ")
			password1 := input()
			fmt.Print("Repeat password: ")
			password2 := input()

			status := register(db, username, password1, password2)
			if status {
				isAuthorized = true
				log.Println("Registration successful")
			} else {
				log.Println("Registration failed")
			}
		case 3:
			return
		default:
			log.Println("Invalid choice. Please select a valid option.")
		}
		if isAuthorized {
			break
		}
	}

	fmt.Printf("Username: %s\n", username)
	menu2()
	for {
		fmt.Print("Choice: ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			notes := getNotes(db, username)
			printNotes(notes)
		case 2:
			fmt.Print("Write note name: ")
			noteName := input()

			fmt.Print("Write note description: ")
			description := input()

			fmt.Print("Note should be private? y/n ")
			choice := input()

			isPrivate := false
			if choice == "y" {
				isPrivate = true
			}

			addNote(db, username, noteName, description, isPrivate)
		case 3:
			fmt.Println("Write note name: ")
			noteName := input()

			deleteNote(db, username, noteName)
		case 4:
			menu2()
		case 5:
			return
		}
	}
}
