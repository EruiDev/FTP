package main

import (
	"fmt"
	"myftp/core"
	"net"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./myftp (port) (path)")
		return
	}
	port := os.Args[1] // TODO: If path is empty take the current directory as the default
	path, err := filepath.Abs(os.Args[2])
	if err != nil {
		fmt.Println("Path does not exist")
		return
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error on listening:", err)
		return
	}
	fmt.Println("Server started on path: ", path)

	defer listener.Close()

	fmt.Println("Server listening on port: ", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error on accepting call: ", err)
			continue
		}
		go core.Protocol(conn, path)
	}
}
