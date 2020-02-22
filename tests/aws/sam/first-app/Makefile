.PHONY: deps clean build debug_prepare clean debug_build debug api_paramter cli debug_api

debug_prepare:
	rm -rf dlv
	mkdir dlv
	GOARCH=amd64 GOOS=linux \
	go build -o dlv github.com/go-delve/delve/cmd/dlv

clean:
	rm -rf ./aws-lambda/aws-lambda

debug_build: | clean
	GOOS=linux GOARCH=amd64 go build -gcflags='-N -l' -o aws-lambda/aws-lambda ./aws-lambda

build: | clean
	GOOS=linux \
	GOARCH=amd64 \
	CGO_ENABLED=0 \
	go build -o aws-lambda/aws-lambda ./aws-lambda

cl_build: clean
	GOOS=linux \
	go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o aws-lambda/aws-lambda ./aws-lambda

api: | build
	sam local start-api --env-vars env_test.json

debug_api: | debug_build
	sam local start-api --env-vars env_test.json \
						-d 5986 \
						--debugger-path dlv \
						--debug-args "-delveAPI=2" \
						--docker-network deveploer_default

api_paramter: | build
	sam local start-api --parameter-overrides SomeVar=define_out_side

cli: | build
	-@docker stop sam_cli
	docker run --rm \
	           --name sam_cli \
	           -p 3000:3000 \
	           -v "$$PWD":/opt/sam \
	           -v /var/run/docker.sock:/var/run/docker.sock \
	           -t sam-cli:v1 \
	           ${PWD}

