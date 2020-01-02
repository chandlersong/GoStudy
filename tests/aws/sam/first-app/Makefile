.PHONY: deps clean build debug_prepare clean debug_build debug api_paramter

.PHONY: deps clean build

debug_prepare:
	rm -rf dlv
	mkdir dlv
	GOARCH=amd64 GOOS=linux
	go build -o dlv github.com/go-delve/delve/cmd/dlv

clean:
	rm -rf ./aws-lambda/aws-lambda

debug_build: | clean
	GOOS=linux GOARCH=amd64 go build -gcflags='-N -l' -o aws-lambda/aws-lambda ./aws-lambda

debug: | debug_build
	sam local start-api -d 5986 --debugger-path dlv --debug-args "-delveAPI=2"

build: | clean
	GOOS=linux GOARCH=amd64 go build -o aws-lambda/aws-lambda ./aws-lambda

api: | build
	sam local start-api --env-vars env_test.json

api_paramter: | build
	sam local start-api --parameter-overrides SomeVar=define_out_side