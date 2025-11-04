.PHONY: all build run gotool clean help

linux:


	CGO_ENABLED=0  GOOS=linux  GOARCH=amd64 go build  -o ./targets/im_sdk_public_server ./main.go
windows:
	CGO_ENABLED=0  GOOS=windows  GOARCH=amd64 go build  -o ./targets/im_sdk_public_server ./main.go
restart:
	sudo supervisorctl restart vi

clean:
	-rm -rf ./targets
