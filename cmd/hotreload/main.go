package main

import (
	"flag"
	"fmt"
	"time"

	"hotreload/internal/builder"
	"hotreload/internal/debounce"
	"hotreload/internal/process"
	"hotreload/internal/watcher"
)

func main() {

	root := flag.String("root", ".", "root directory to watch")
	build := flag.String("build", "", "build command")
	run := flag.String("exec", "", "run command")

	flag.Parse()

	changes := make(chan struct{}, 1)

	debounced := debounce.Debounce(500*time.Millisecond, changes)

	server := &process.Server{}

	err := watcher.Watch(*root, changes)
	if err != nil {
		fmt.Println("Watcher error:", err)
		return
	}

	fmt.Println("Starting hotreload...")

	fmt.Println("Running initial build")

	err = builder.Build(*build)

	if err != nil {
		fmt.Println("Initial build failed")
	} else {
		server.Start(*run)
	}

	for range debounced {

		fmt.Println("File change detected. Rebuilding...")

		server.Stop()

		err := builder.Build(*build)

		if err != nil {
			fmt.Println("Build failed. Waiting for next change...")
			continue
		}

		server.Start(*run)
	}
}