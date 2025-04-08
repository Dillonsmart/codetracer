package main

import (
	"bufio"
	"fmt"
	"os"
)

var W *Watcher

func StartProg() {
	// Init a new watcher
	W := CreateWatcher()

	homeDir, _ := os.UserHomeDir()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the directory to watch: " + homeDir + "/")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	W.SetIgnoreHidden(true)
	W.SetDir(homeDir + "/" + input)
	W.Start()
}
