build:
	go build -o hotreload ./cmd/hotreload

server:
	go build -o ./bin/server ./testserver

run:
	./hotreload --root ./testserver --build "go build -o ./bin/server ./testserver" --exec "./bin/server"

demo: build server run

test:
	go test ./...