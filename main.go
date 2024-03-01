package main

import (
	"fmt"

	goblincli "github.com/chturner94/goblin-logger/pkg/GoblinCLI"
)

func main() {
	app := goblincli.App{
		MenuOptions: []goblincli.MenuOptions{
			{EventFunction: option1, MenuEntry: "1. Configure a new job"},
			{EventFunction: option2, MenuEntry: "2. Get help"},
		},
	}
	app.Run()
}

func option1() {
	fmt.Println("You chose option1")
}

func option2() {
	fmt.Println("You chose option2")
}
