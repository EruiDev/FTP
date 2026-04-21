package core

import (
	"bufio"
	"fmt"
	"io"
	"myftp/commands"
	"myftp/commons"
	"net"
	"path/filepath"
	"strings"
	"time"
)

func Protocol(conn net.Conn, path string) {
	defer conn.Close()
	defer fmt.Println("Client disconnected:", conn.RemoteAddr())
	cleanPath := filepath.Clean(path)
	user := commons.Info{
		OriginalDir:    cleanPath,
		CurrentDir:     cleanPath,
		Username:       "",
		IsLogged:       false,
		Conn:           conn,
		DataConnection: nil,
	}
	defer func() {
		if user.DataConnection != nil {
			user.DataConnection.Close()
		}
	}()
	fmt.Println("New client connected:", conn.RemoteAddr())
	conn.Write([]byte(commons.Welcome))
	reader := bufio.NewReader(conn)
	protocolLoop(conn, &user, reader)
}

func protocolLoop(conn net.Conn, user *commons.Info, reader *bufio.Reader) {
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Minute))
		split, err := getParsedCommand(reader)
		if err != nil {
			return
		}
		if len(split) == 0 || split[0] == "" {
			continue
		}
		split[0] = strings.ToUpper(split[0])

		if split[0] == "PASS" {
			fmt.Println("Command received: [PASS ****]")
		} else {
			fmt.Println("Command received:", split)
		}

		if split[0] == "QUIT" {
			conn.Write([]byte(commons.Goodbye))
			if user.DataConnection != nil {
				user.DataConnection.Close()
				user.DataConnection = nil
			}
			break
		}
		command, ok := commandList[split[0]]
		if !ok {
			conn.Write([]byte(commons.InvalidCommand))
			continue
		}
		if !user.IsLogged && (split[0] != "USER" && split[0] != "PASS" && split[0] != "HELP" && split[0] != "SYST" && split[0] != "FEAT") {
			conn.Write([]byte(commons.LoginFirst))
			continue
		}

		if split[0] == "HELP" {
			if err := commands.HandleHelp(split, user, commandList); err != nil {
				conn.Write([]byte(commons.InternalError))
				fmt.Println("Internal error: ", err)
			}
			continue
		}

		if err := command.Handler(split, user); err != nil {
			conn.Write([]byte(commons.InternalError))
			fmt.Println("Internal error: ", err)
			continue
		}
	}
}

func getParsedCommand(reader *bufio.Reader) ([]string, error) {
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

	idx := strings.IndexByte(line, ' ')
	if idx == -1 {
		return []string{line}, nil
	}
	return []string{line[:idx], line[idx+1:]}, nil
}
