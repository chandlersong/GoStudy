.PHONY: deps clean build debug_prepare clean debug_build debug

.PHONY: deps clean build

debug_prepare:
	rm -rf dlv
	mkdir dlv
	GOARCH=amd64 GOOS=linux
	go build -o dlv github.com/go-delve/delve/cmd/dlv

clean:
	rm -rf ./hello-world/hello-world

debug_build: | clean
	GOOS=linux GOARCH=amd64 go build -gcflags='-N -l' -o hello-world/hello-world ./hello-world

debug: | debug_build
	sam local start-api -d 5986 --debugger-path dlv --debug-args "-delveAPI=2"

build: | clean
	GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world

api: | build
	sam local start-api