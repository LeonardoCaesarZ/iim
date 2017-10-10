package main

import "iim/frame/tcp"

import "fmt"

func main() {
	server := tcp.NewServer(9998, handler)
	server.Loop()
}

func handler(context *tcp.Context) interface{} {
	fmt.Println(context)
	return "bbb"
}
