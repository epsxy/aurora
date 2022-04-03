.PHONY: release-darwin-amd64, release-windows-amd64

run:
	go run main.go

build:
	go build -o ./bin/aurora main.go

install:
	go build -o ./bin/aurora main.go
	sudo cp ./bin/aurora /usr/local/bin/aurora

uninstall:
	sudo rm -rf /usr/local/bin/aurora

install-git:
	sudo cp ./release/git-aur /usr/local/bin/git-aur
	sudo chmod +x /usr/local/bin/git-aur

uninstall-git:
	sudo cp ./release/git-aur /usr/local/bin/git-aur
	sudo chmod +x /usr/local/bin/git-aur

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64 main.go &

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin-amd64 main.go &

build-freebsd-amd64:
	GOOS=freebsd GOARCH=amd64 go build -o bin/freebsd-amd64 main.go &

build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -o bin/windows-amd64 main.go &

release-darwin-amd64: build-darwin-amd64
	mkdir -p release/darwin-amd64
	cp release/git-aur release/darwin-amd64
	cp release/Makefile release/darwin-amd64
	cp release/README.md release/darwin-amd64
	cp bin/darwin-amd64 release/darwin-amd64/aurora
	tar -czvf release/darwin-amd64.tar.gz release/darwin-amd64
	rm -rf release/darwin-amd64