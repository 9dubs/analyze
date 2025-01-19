Analyze, a simple CLI in Go, to analyze text files.

make sure to have Go install, verify using `go --version`.

Clone the repo and verify the import path for `rootcmd.go` properly. alternatively, run `go mod tidy`.

run the main go file along with a sample file to test the code.
```
go run main.go sample.txt
```

building an executable using `go build` and placing this executable inside `/usr/bin` will run `analyze` as a `shell` command.
