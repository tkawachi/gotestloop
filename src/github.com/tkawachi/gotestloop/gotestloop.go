package main

import (
	"github.com/howeyc/fsnotify"
	"log"
//	"bytes"
	"os/exec"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("start watching .")
	err = watcher.Watch(".")
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case ev := <-watcher.Event:
			log.Println("event:", ev)
			cmd := exec.Command("go", "test")
			out, err := cmd.CombinedOutput()
			log.Print(string(out))
			if err != nil {
				log.Println(err)
			}
		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}
}
