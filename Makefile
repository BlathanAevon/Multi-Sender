compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=386 go build -o build/linux/sender cmd/MultiSender/main.go
	GOOS=darwin GOARCH=amd64 go build -o build/mac/sender cmd/MultiSender/main.go
	GOOS=windows GOARCH=386 go build -o build/windows/sender cmd/MultiSender/main.go
