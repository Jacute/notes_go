package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var reader = bufio.NewReader(os.Stdin)

func input() string {
	inputString, _ := reader.ReadString('\n')
	inputString = strings.TrimSpace(inputString)

	return inputString
}

func getHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

func compareHash(password string, storedHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
