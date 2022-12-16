package main

import (
	"fmt"
	"goMokeney/repl"
	"os"
	"os/user"
)

func main(){
	user, err := user.Current()
	if err != nil{
		panic(err)
	}
	fmt.Printf("Hello %s! This is the monkey programming language\n", user.Username)
	fmt.Printf("Feel free to commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
