all: format lint test install

format:
	go fmt ./...

lint:
	golint `go list ./...`

generate:
	go generate ./...

swagger: validate
	swagger generate server -t api/swagger --spec=specification/swagger.json --exclude-main -A hello

validate: specification/swagger.json
	swagger validate $<

specification/swagger.json: specification/*.yaml
	swagger flatten --output=$@ specification/api.yaml

test:
	go test ./...

install:
	go install -v ./...

clean:
	go clean ./...
