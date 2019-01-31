.PHONY: deps clean build

deps:
	go get -u github.com/aws/aws-lambda-go/events
	go get -u github.com/aws/aws-lambda-go/lambda
	go get -u github.com/holiday-jp/holiday_jp-go

clean: 
	rm -rf ./AutoStopStart-EC2/AutoStopStart-EC2
	
build:
	GOOS=linux GOARCH=amd64 go build -o AutoStopStart-EC2/AutoStopStart-EC2 ./AutoStopStart-EC2

