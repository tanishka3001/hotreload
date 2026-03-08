# HotReload – Go Development Hot Reload Engine

## Overview

HotReload is a lightweight CLI tool that automatically rebuilds and restarts a Go server whenever source files change. It is designed to improve developer productivity by eliminating the need to manually stop, rebuild, and restart a server after every code modification.

The tool monitors a project directory, triggers a rebuild when relevant files change, and restarts the running server with the updated code.

---

## Features

### Automatic File Watching

The tool monitors the project directory for file system events using `fsnotify`. Changes to relevant source files trigger automatic rebuilds and restarts.

### Debounced Rebuilds

Modern editors often generate multiple filesystem events when saving a file. To prevent redundant builds, HotReload implements a debounce mechanism that groups rapid file events into a single rebuild trigger.

### Automatic Server Restart

When a change is detected:

1. The currently running server process is terminated.
2. The project is rebuilt using the provided build command.
3. The server is restarted using the provided execution command.

### Initial Build on Startup

HotReload performs a build immediately on startup so developers do not need to manually trigger the first rebuild.

### Real-Time Log Streaming

Server logs are streamed directly to the terminal, allowing developers to see output in real time without buffering.

### Recursive Directory Watching

All directories within the project are monitored recursively. Nested folders are automatically included during startup.

### Dynamic Directory Detection

If new directories are created while the tool is running, they are automatically added to the watcher.

### Graceful Handling of Deleted Folders

When directories are deleted, the watcher handles the event safely without crashing.

### File Filtering

HotReload ignores directories and files that should not trigger rebuilds, including:

* `.git/`
* `node_modules/`
* `bin/`
* temporary editor files (`.swp`, `.tmp`, backup files)

### Crash Protection

A restart cooldown prevents rapid restart loops if the server crashes immediately after launch.

---

## CLI Usage

The CLI follows the format:

```
hotreload --root <project-folder> --build "<build-command>" --exec "<run-command>"
```

### Example

```
hotreload \
  --root ./testserver \
  --build "go build -o ./bin/server ./testserver" \
  --exec "./bin/server"
```

Parameters:

| Flag      | Description                         |
| --------- | ----------------------------------- |
| `--root`  | Project directory to watch          |
| `--build` | Command used to rebuild the project |
| `--exec`  | Command used to start the server    |

---

## Project Structure

```
hotreload/
│
├── cmd/
│   └── hotreload/
│       └── main.go
│
├── internal/
│   ├── builder/
│   │   └── builder.go
│   ├── process/
│   │   └── process.go
│   ├── watcher/
│   │   └── watcher.go
│   └── debounce/
│       └── debounce.go
│
├── testserver/
│   └── main.go
│
├── bin/
├── runserver.ps1
└── README.md
```

### Architecture

**Watcher**

* Monitors filesystem events using `fsnotify`
* Registers watchers recursively across project directories
* Handles creation and deletion of directories dynamically

**Debounce**

* Aggregates rapid file events
* Prevents redundant rebuilds caused by editor save behavior

**Builder**

* Executes the provided build command
* Ensures the latest code is compiled before restarting the server

**Process Manager**

* Handles server lifecycle
* Terminates existing server processes and child processes before restart

---

## Running the Demo

### 1. Build the HotReload CLI

```
go build -o hotreload ./cmd/hotreload
```

### 2. Build the Example Server

```
go build -o ./bin/server ./testserver
```

### 3. Run HotReload

```
hotreload \
  --root ./testserver \
  --build "go build -o ./bin/server ./testserver" \
  --exec "./bin/server"
```

### 4. Modify Code

Edit any file in the `testserver` directory and save it.

HotReload will automatically:

* detect the change
* rebuild the project
* restart the server

---

## Testing

A unit test is included for the debounce component to validate event aggregation behavior.

Run tests with:

```
go test ./...
```

---

## Scalability Considerations

To avoid hitting operating system limits on file watchers:

* watchers are registered per directory rather than per file
* ignored directories reduce unnecessary watchers
* recursive directory walking registers watchers efficiently

---

## Demo Server

A simple HTTP server is included in the `testserver` directory to demonstrate hot reload functionality.

The server responds at:

```
http://localhost:8080
```

---

