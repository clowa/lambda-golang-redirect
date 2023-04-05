.PHONY: build clean deploy

build:
	go mod verify
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/lambda-golang-redirect ./functions/redirect/main.go

clean:
	rm -rf ./bin
#	serverless delete_domain
#	serverless remove-cert

deploy: clean build
	serverless deploy --verbose
