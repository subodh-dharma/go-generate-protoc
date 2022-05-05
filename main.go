package main

import "fmt"

func main() {
	fmt.Println("hello! I am main.")
}

//go:generate protoc --go_out=. --go-grpc_out=. -I=./pb ./pb/employee.proto
