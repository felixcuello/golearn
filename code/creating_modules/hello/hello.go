package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Manga de trolos")
	fmt.Println(message)
}
