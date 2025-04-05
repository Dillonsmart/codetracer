package main

import (
	"fmt"
	"time"
)

var StartTime string
var W *Watcher

func StartProg() {
	StartTime = time.Now().Format("2006-01-02 15:04:05")

	// Init a new watcher
	W := CreateWatcher()
	W.Start()

	fmt.Println("Codetracer started at", W.FormattedStartTime())

}
