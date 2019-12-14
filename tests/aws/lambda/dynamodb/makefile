BINARY="example"
VERSION=1.0.0
BUILD=`date +%FT%T%z`

export GOOS=linux
export GO111MODULE=on

default:
	@echo hello world of go lambda

listEnv:
	@echo ${GOOS}
	@echo ${USERPROFILE}
	
clean:
	@echo clean package
	@rm -rf build
	@mkdir build

build: | clean
	@echo start to build
	@go build -o build/main main.go
	
lambdaPackage: | build
	@echo start to package lambda to main.zip
	@${USERPROFILE}/Go/bin/build-lambda-zip.exe build/main main.zip
	@mv main.zip build/

updateLambda: | lambdaPackage
	@echo start to update lambda code
	@aws lambda  update-function-code \
	--function-name my-function \
	--zip-file fileb://build/main.zip



.PHONY: default listEnv build lambdaPackage clean updateLambda