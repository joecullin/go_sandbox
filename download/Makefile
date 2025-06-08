run:
	# clean dir:
	rm -f files/*

	# run program:
	go run download.go

	# check results:
	ls -halt files

format:
	gofmt -w *.go

build:
	GOOS=darwin GOARCH=amd64 go build -o bin/download-amd64-mac download.go
	ls -halt bin

replace-test:
	# clean dir:
	rm -f files/*

	# run program from bin dir:
	./bin/download-amd64-mac

	# check results:
	ls -halt bin
