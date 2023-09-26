package main

import (
	"example.com/greetings"
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
