package main

import "time"

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

func (w *Watcher) Start() {
	w.StartTime = time.Now()
	w.running = true
}

func (w *Watcher) Stop() {
	w.StopTime = time.Now()
	w.running = false
}

func (w *Watcher) SetIgnoreHidden(v bool) {
	w.ignoreHidden = v
}

func (w *Watcher) SetDir(d string) {
	w.dir = d
}

func (w *Watcher) FormattedStartTime() string {
	return w.StartTime.Format("2006-01-02 15:04:05")
}
