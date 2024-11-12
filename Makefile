build:
	@go build 

install:
	@go build 
	@go install 

clean:
	@rm fileshare-relay

run:
	@go build 
	@go install 
	@fileshare-relay
