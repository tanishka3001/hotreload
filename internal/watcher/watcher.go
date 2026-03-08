package watcher

import (
	"fmt"
	"strings"
	"github.com/fsnotify/fsnotify"
	"path/filepath"
	"os"
)
func shouldIgnore(path string) bool {

	ignoreDirs := []string{
		".git",
		"node_modules",
		"bin",
	}

	for _, dir := range ignoreDirs {
		if strings.Contains(path, dir) {
			return true
		}
	}

	if strings.HasSuffix(path, "~") ||
		strings.HasSuffix(path, ".tmp") ||
		strings.HasSuffix(path, ".swp") {
		return true
	}

	return false
}
func addRecursiveWatch(w *fsnotify.Watcher, root string) error {

	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			if shouldIgnore(path) {
				return filepath.SkipDir
			}

			fmt.Println("Watching directory:", path)

			return w.Add(path)
		}

		return nil
	})
}
func Watch(path string, changes chan struct{}) error {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	err = addRecursiveWatch(watcher, path)
if err != nil {
	return err
}

	go func() {
		for {
			select {

			case event := <-watcher.Events:
				if shouldIgnore(event.Name) {
	              return
}

fmt.Println("File change detected:", event.Name)
changes <- struct{}{}

			case err := <-watcher.Errors:
				fmt.Println("Watcher error:", err)
			}
		}
	}()

	return nil
}