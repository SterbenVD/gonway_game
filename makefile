bin_file = ./bin/exec

run: build
	./bin/exec

build:
	go build -o $(bin_file)  ./main.go

clean:
	rm -rf ./bin