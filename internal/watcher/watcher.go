package watcher

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func Watch(path string, changes chan struct{}) error {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	err = watcher.Add(path)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {

			case event := <-watcher.Events:
				fmt.Println("File change detected:", event.Name)
				changes <- struct{}{}

			case err := <-watcher.Errors:
				fmt.Println("Watcher error:", err)
			}
		}
	}()

	return nil
}