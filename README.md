# go-generate-protoc

This is a example on how to use the `//go:generate` compiler directive to generate go grpc protobuf code.
This introduces a documented way to generate protobuf go source code and can be plugged in build pipeline very easily.

# Pre-requisite

1. Install the following protocol buffer compilers and grpc code generator
```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

2. Update your path

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

3. Generate the grpc and protobuf files.

```
$ go generate ./...
```

That's it! The grpc service and protobuf files will be generated in the configured package.


# How it works!

The main magic here lies underneath the `//go:generate` comment which I have declared in `main.go`. 

```go
//go:generate protoc --go_out=. --go-grpc_out=. -I=./pb ./pb/employee.proto
```

The above go compiler directive can be declared anywhere in this go package, although you will need keep the relative paths for the options that you may pass in the command.

The `go:generate` directive can be broken down as

`//go:generate <command> --options <target>`

Here the `<command>` can be any command that you want to run during compile time. 
In our case we will be running `protoc` with options. Ensure that `protoc` is configured in your PATH.
The rest of the command for `protoc` will be same as if you are running it from the root directory of this repo.

NOTE: One important thing to consider is where this `//go:generate` directive is declared. Relative to the file in which it is declared, your `protoc` command may vary.

For e.g. if you decide to declare the go directive in `employee/employee.go` then the comment will look like following

```go
//go:generate protoc --go_out=./.. --go-grpc_out=./.. -I=../pb ../pb/employee.proto
```

IMHO, it is better to keep all the //go:generate directive in a common file like main.go so that it is easy to manage all at one place.
