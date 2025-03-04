package core

import (
	"bufio"
	"fmt"
	"io"
	"myftp/commons"
	"net"
	"strings"
)

func Protocol(conn net.Conn, path string) {
	defer conn.Close()
	defer fmt.Println("Client disconnected: ", conn.RemoteAddr())
	user := commons.Info{
		CurrentDir:     path,
		Username:       "",
		IsLogged:       false,
		Conn:           conn,
		DataConnection: nil,
	}
	fmt.Println("New client connected: ", conn.RemoteAddr())
	conn.Write([]byte("220 Welcome to MyFTP\r\n"))
	protocol_loop(conn, &user)
}

func protocol_loop(conn net.Conn, user *commons.Info) {
	for {
		split, err := getParsedCommand(conn)
		fmt.Println("Command received: ", split)
		if err != nil {
			return
		}
		if len(split) == 0 || split[0] == "" {
			continue
		}
		split[0] = strings.ToUpper(split[0])
		if split[0] == "QUIT" {
			conn.Write([]byte(commons.Goodbye))
			break
		}
		command, ok := commandList[split[0]]
		if !ok {
			conn.Write([]byte(commons.InvalidCommand))
			continue
		}
		if !user.IsLogged && (split[0] != "USER" && split[0] != "PASS") {
			conn.Write([]byte(commons.LoginFirst))
			continue
		}
		if err := command.Handler(split, user); err != nil {
			conn.Write([]byte(commons.InternalError))
			fmt.Println("Internal error: ", err)
			return
		}
	}
}

func getParsedCommand(conn net.Conn) ([]string, error) {
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, err
	}

	line = strings.TrimSpace(line)
	if line == "" {
		return nil, nil
	}

	split := strings.Split(line, " ")
	return split, nil
}
