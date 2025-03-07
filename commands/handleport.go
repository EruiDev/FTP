package commands

import (
	"myftp/commons"
	"myftp/utils"
	"net"
	"strconv"
	"strings"
)

func HandlePort(args []string, info *commons.Info) error {
	if !info.IsLogged {
		return utils.WriteMessage(info.Conn, commons.LoginFirst)
	}

	if len(args) != 2 {
		return utils.WriteMessage(info.Conn, commons.InvalidCommand)
	}

	parts := strings.Split(args[1], ",")
	if len(parts) != 6 {
		return utils.WriteMessage(info.Conn, "501 Syntax error in parameters.\r\n")
	}

	ip := make([]string, 4)
	for i := 0; i < 4; i++ {
		val, err := strconv.Atoi(parts[i])
		if err != nil || val < 0 || val > 255 {
			return utils.WriteMessage(info.Conn, "501 Syntax error in parameters.\r\n")
		}
		ip[i] = parts[i]
	}

	p1, err := strconv.Atoi(parts[4])
	if err != nil || p1 < 0 || p1 > 255 {
		return utils.WriteMessage(info.Conn, "501 Syntax error in parameters.\r\n")
	}

	p2, err := strconv.Atoi(parts[5])
	if err != nil || p2 < 0 || p2 > 255 {
		return utils.WriteMessage(info.Conn, "501 Syntax error in parameters.\r\n")
	}

	port := p1*256 + p2

	address := strings.Join(ip, ".") + ":" + strconv.Itoa(port)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return utils.WriteMessage(info.Conn, "425 Can't open data connection.\r\n")
	}

	info.DataConnection = conn
	return utils.WriteMessage(info.Conn, "200 PORT command successful.\r\n")
}
