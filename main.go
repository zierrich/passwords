package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	passCount    = 10
	passLen      = 18
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
	special      = "!@#$%&*:;[]{}?"
	letters      = lowerLetters + upperLetters
	allChars     = letters + digits + special
)

func genPassword() string {
	p := make([]byte, 0, passLen)
	p = append(p, randChar(letters))
	letterRun := 1
	digitRun := 0
	specialCount := 0
	lastType := "letter"

	for len(p) < passLen {
		pool := buildPool(letterRun, digitRun, specialCount, lastType)
		c := randChar(pool)
		currType := charType(c)

		switch currType {
		case "letter":
			letterRun++
			digitRun = 0
		case "digit":
			digitRun++
			letterRun = 0
		case "special":
			specialCount++
			letterRun = 0
			digitRun = 0
		}

		p = append(p, c)
		lastType = currType
	}

	return string(p)
}

func buildPool(letterRun, digitRun, specialCount int, lastType string) string {
	pool := ""

	if letterRun < 3 {
		pool += letters
	}

	if digitRun < 3 {
		pool += digits
	}

	if specialCount < 4 && lastType != "special" {
		pool += special
	}

	return pool
}

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

func randChar(s string) byte {
	return s[rand.Intn(len(s))]
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < passCount; i++ {
		fmt.Println(genPassword())
	}
}