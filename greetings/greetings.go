package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	return randomFormat(), nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi,%v.Welcome",
		"Great to see you,%v!",
		"Hail,%v!Well met!",
	}
	fmt.Println(rand.Intn(len(formats)))

	return formats[rand.Intn(len(formats))]
}
