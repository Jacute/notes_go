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

func printNotes(notes []Note) {
	if len(notes) == 0 {
		fmt.Println("No notes :(\n======")
		return
	}

	for index, note := range notes {
		fmt.Printf(yellow("%d\n")+"Author: %s\n"+blue("=======\n")+"Note: %s\nDescription: %s\n"+blue("=======\n"), index+1, note.author, note.name, note.description)
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
				log.Println(green("Login successful"))
			} else {
				log.Println(red("Login failed"))
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
				log.Println(green("Registration successful"))
			} else {
				log.Println(red("Registration failed"))
			}
		case 3:
			return
		default:
			log.Println(red("Invalid choice. Please select a valid option."))
		}
		if isAuthorized {
			break
		}
	}

	fmt.Printf("Username: %s\n", cyan(username))
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
			noteDescription := input()

			fmt.Print("Note should be private? y/n ")
			choice := input()

			is_private := false
			if choice == "y" {
				is_private = true
			}

			note := Note{
				name:        noteName,
				description: noteDescription,
				is_private:  is_private,
				author:      username,
			}

			addNote(db, note)
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
