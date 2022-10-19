package main

import (
	"fmt"
	"time"
	"github.com/tcnksm/go-input"
	"os"
)

func main() {

	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}
	
	query := "What is your name?"
	_name, err := ui.Ask(query, &input.Options{
		Default: "tcnksm",
		Required: true,
		Loop:     true,
	})
	fmt.Print(_name,err)
	go sayHello()
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}
