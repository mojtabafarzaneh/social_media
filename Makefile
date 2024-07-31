build:
	@go build -o bin/api

help: build
	@./bin/api

serve: 
	@go run main.go serve

work:
	@CompileDaemon -command=./[executable-file]
