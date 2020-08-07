# go build command
build:
	@echo " >> building binaries"
	@go build -v -o bin/rest cmd/rest/*.go

# go run command
run: build
	./url-shortener-service

# run all go:generate commands (eg. Mock files generator)
generate:
	@go generate ./...