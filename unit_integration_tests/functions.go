package main

import (
	"regexp"
	"unicode"
)

// unit test for use case 1
// for use case 1
func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func isPassValid(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

// for use case 4: Create Room
// type room struct {
// 	Id          int
// 	Name        string
// 	PigstyNum   int
// 	Humidity    string
// 	Temperature string
// }

// func createRoom() {

// 	room := room{
// 		Id:          1,
// 		Name:        "Room 1",
// 		PigstyNum:   1,
// 		Humidity:    "30% Humid",
// 		Temperature: "72F",
// 	}

// }
