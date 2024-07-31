build:
	@go build -o bin/api

help: build
	@./bin/api

serve: 
	@go run main.go serve

migrate:
	@go run main.go migrate

seed:
	@go run main.go seed

work:
	@CompileDaemon -command=./[executable-file]
