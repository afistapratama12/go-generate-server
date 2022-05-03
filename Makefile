build:
	go build
	rm ~/go-gen-server/go-gen
	mv ./go-gen-server ~/go-gen-server/go-gen

linux: 
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o grader-cli-linux-amd64 . ; \
	cd - >/dev/null

darwin:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o grader-cli-darwin-amd64 . ; \
	cd - >/dev/null

windows:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o grader-cli-windows-amd64.exe . ; \
	cd - >/dev/null
