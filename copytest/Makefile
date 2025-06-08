.SILENT: run
run:
	echo
	echo "############## test copy text file"

	go run copytest.go --copytest

	# check results:
	ls -halt files

	echo
	echo "############## test replace self (ctrl-c to stop)"

	@echo "before:"
	./bin/copytest --version

	./bin/copytest --replaceself

format:
	gofmt -w *.go

build:
	perl -pi -e 's/Version=./Version=1/g' copytest.go
	GOOS=darwin GOARCH=amd64 go build -o bin/copytestV1 copytest.go
	GOOS=windows GOARCH=amd64 go build -o bin/copytestV1.exe copytest.go
	perl -pi -e 's/Version=./Version=2/g' copytest.go
	GOOS=darwin GOARCH=amd64 go build -o bin/copytestV2 copytest.go
	GOOS=windows GOARCH=amd64 go build -o bin/copytestV2.exe copytest.go
	cp bin/copytestV1 bin/copytest
	cp bin/copytestV1.exe bin/copytest.exe
	ls -halt bin/copytest*
	./bin/copytest --version
	./bin/copytestV1 --version
	./bin/copytestV2 --version
	@echo

wincopy:
	cp bin/*.exe wintest
	rm -f wintest/wintest.zip
	pushd wintest && zip wintest.zip *.exe && popd

local-replace-test:
	make build

	# run program from bin dir:
	./bin/copytest-amd64-mac

	# check results:
	ls -halt files
