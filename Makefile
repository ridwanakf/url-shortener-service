# go build command
build:
	@echo " >> building binaries"
	@go build -v -o url-shortener-service cmd/*.go

# go run command
run: build
	./url-shortener-service

# run all go:generate commands (eg. Mock files generator)
generate:
	@go generate ./...