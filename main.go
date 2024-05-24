package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is Monkey!\n", user.Username)
	fmt.Println("Type some stuff")
	repl.Start(os.Stdin, os.Stdout)
}
