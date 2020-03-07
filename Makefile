.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/timer

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy --verbose
