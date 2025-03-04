package main

import (
	"fmt"
	"myftp/core"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./my_ftp (port) (path)")
		return
	}
	port := os.Args[1]
	path := os.Args[2]

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Path does not exist")
		return
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error on listening:", err)
		return
	}

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
