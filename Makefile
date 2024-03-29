.PHONY: build clean deploy

build:
	go mod verify
	env GOARCH=arm64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o bootstrap ./functions/redirect/main.go

clean:
	rm -rf ./bin
#	serverless delete_domain
#	serverless remove-cert

deploy: clean build
	serverless deploy --verbose
