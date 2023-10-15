env go build -o ./dist/ElaParse-linux main.go &&
env GOOS=darwin GOARCH=amd64 go build -o ./dist/ElaParse-macos main.go &&
env GOOS=windows GOARCH=amd64 go build -o ./dist/ElaParse-win-64.exe main.go &&
env GOOS=windows GOARCH=386 go build -o ./dist/ElaParse-win-386.exe main.go
