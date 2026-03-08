package main

import (
	"flag"
	"time"
	"log/slog"
	"os"

	"hotreload/internal/builder"
	"hotreload/internal/debounce"
	"hotreload/internal/process"
	"hotreload/internal/watcher"
)

func main() {
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	root := flag.String("root", ".", "root directory to watch")
	build := flag.String("build", "", "build command")
	run := flag.String("exec", "", "run command")

	flag.Parse()

	changes := make(chan struct{}, 1)

	debounced := debounce.Debounce(500*time.Millisecond, changes)

	server := &process.Server{}

	err := watcher.Watch(*root, changes)
	if err != nil {
		logger.Error("Watcher error:", "error", err)
		return
	}

	logger.Info("Starting hotreload...")

	logger.Info("Running initial build")

	err = builder.Build(*build)

	if err != nil {
		logger.Error("Initial build failed", "error", err)
	} else {
		server.Start(*run)
	}

	for range debounced {

		logger.Info("File change detected. Rebuilding...")

		server.Stop()
		err := builder.Build(*build)

if err != nil {
    logger.Error("Build failed", "error", err)
    continue
}

time.Sleep(1 * time.Second)

logger.Info("Restarting server")

server.Start(*run)

	}
}