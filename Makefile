build:
	@go build -o bin/api

run: build
	@./bin/api

work:
	@CompileDaemon -command=./[executable-file]
