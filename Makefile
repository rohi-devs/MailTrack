build:
	go build -o bin/cli .
run : build
	./bin/cli