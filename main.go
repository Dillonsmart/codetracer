package main

import (
	"codetracer/utils"
	"fmt"
	"os"
	"strings"
)

// defines the allowed arguments
const (
	Start utils.AllowedString = "start"
	Stop  utils.AllowedString = "stop"
)

func main() {
	commandArguments := os.Args[1:]

	if len(commandArguments) == 0 {
		fmt.Println("Please pass an argument")
		return
	}

	action := strings.ToLower(commandArguments[0])

	println(utils.IsStringAllowed([]utils.AllowedString{Start, Stop}, action))
}
