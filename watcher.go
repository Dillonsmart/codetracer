package main

import (
	"fmt"
	"time"
)

type Watcher struct {
	StartTime    time.Time
	StopTime     time.Time
	running      bool
	ignoreHidden bool
	dir          string
}

func CreateWatcher() *Watcher {
	return &Watcher{
		running:      false,
		ignoreHidden: true,
	}
}

func (w *Watcher) Start() *Watcher {
	w.StartTime = time.Now()
	w.running = true

	printSpacer(w.FormattedStartTime())
	fmt.Println("Codetracer started")
	fmt.Println("Watching directory:", w.dir)

	return w
}

func (w *Watcher) Stop() *Watcher {
	w.StopTime = time.Now()
	w.running = false

	return w
}

func (w *Watcher) SetIgnoreHidden(v bool) *Watcher {
	w.ignoreHidden = v

	return w
}

func (w *Watcher) SetDir(d string) *Watcher {
	w.dir = d

	return w
}

func (w *Watcher) FormattedStartTime() string {
	return w.StartTime.Format("2006-01-02 15:04:05")
}

func printSpacer(prefix string) {
	fmt.Println("")
	fmt.Println(prefix + " ===================================================")
}
