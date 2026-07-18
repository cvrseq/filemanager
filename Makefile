
run:
	@echo "running binary ... "
	go run ./cmd/main.go

build:
	@echo "compiling project ..."
	mkdir -p ./bin
	go build -o ./bin/app ./cmd/main.go

clean:
	@echo "deleting binaries"
	rm -rf ./bin
	
