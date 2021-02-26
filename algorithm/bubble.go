package algorithm

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func bubble() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
