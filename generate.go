package main

import (
	"math/rand"
	"time"
)

const (
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
	special      = "!@#$%&*:;[]{}?"
	letters      = lowerLetters + upperLetters
	allChars     = letters + digits + special
)

// Generate creates a random password of the specified length
// Rules:
// - First character is always a letter
// - Letters: 1–3 consecutive allowed
// - Digits: 1–2 consecutive allowed
// - Specials: exactly 4 in total, never consecutive
func Generate(length int) string {
	rand.Seed(time.Now().UnixNano()) // seed the random generator each call

	p := make([]byte, 0, length)     // password buffer
	p = append(p, randChar(letters)) // start with a letter
	letterRun, digitRun, specialCount := 1, 0, 0
	lastType := "letter"

	for len(p) < length {
		pool := buildPool(letterRun, digitRun, specialCount, lastType)
		c := randChar(pool)
		currType := charType(c)

		// disallow consecutive special characters
		if currType == "special" && lastType == "special" {
			continue
		}

		// enforce limits on consecutive characters
		switch currType {
		case "letter":
			if letterRun >= 3 {
				continue
			}
			letterRun++
			digitRun = 0
		case "digit":
			if digitRun >= 2 {
				continue
			}
			digitRun++
			letterRun = 0
		case "special":
			if specialCount >= 4 {
				continue
			}
			specialCount++
			letterRun, digitRun = 0, 0
		}

		p = append(p, c)
		lastType = currType
	}

	// ensure exactly 4 special characters
	for i := 1; i < len(p) && specialCount < 4; i++ {
		if charType(p[i]) != "special" &&
			(i == 0 || charType(p[i-1]) != "special") &&
			(i == len(p)-1 || charType(p[i+1]) != "special") {
			p[i] = randChar(special)
			specialCount++
		}
	}

	return string(p)
}

// buildPool determines which characters can be chosen next
// based on consecutive counts and last character type
func buildPool(letterRun, digitRun, specialCount int, lastType string) string {
	pool := ""
	if letterRun < 3 {
		pool += letters
	}
	if digitRun < 2 {
		pool += digits
	}
	if specialCount < 4 && lastType != "special" {
		pool += special
	}
	return pool
}

// charType classifies a character as "letter", "digit", or "special"
func charType(b byte) string {
	switch {
	case isLetter(b):
		return "letter"
	case isDigit(b):
		return "digit"
	default:
		return "special"
	}
}

// randChar returns a random character from the given string
func randChar(s string) byte {
	return s[rand.Intn(len(s))]
}

// isLetter checks if the byte is an English letter
func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

// isDigit checks if the byte is a digit
func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
