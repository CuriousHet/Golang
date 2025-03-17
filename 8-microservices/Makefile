build:
	go build -o catfact-service main.go api.go logging.go service.go types.go

run: build
	./catfact-service

fmt:
	go fmt ./...

clean:
	rm -f catfact-service
