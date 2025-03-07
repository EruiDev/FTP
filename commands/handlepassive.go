package commands

import (
	"math/rand/v2"
	"myftp/commons"
	"myftp/utils"
	"net"
	"strconv"
)

func HandlePassive(args []string, info *commons.Info) error {
	if !info.IsLogged {
		return utils.WriteMessage(info.Conn, commons.LoginFirst)
	}

	var listener net.Listener
	var err error
	var port int
	maxRetries := 10

	for i := 0; i < maxRetries; i++ {
		port = rand.IntN(65535-1024) + 1024
		listener, err = net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port))
		if err == nil {
			break
		}
	}

	if err != nil {
		return utils.WriteMessage(info.Conn, "425 Can't open data connection.\r\n")
	}
	defer listener.Close()

	utils.WriteMessage(info.Conn, "227 Entering Passive Mode (127,0,0,1,"+strconv.Itoa(port>>8)+","+strconv.Itoa(port&0xff)+")\r\n")
	conn, err := listener.Accept()
	if err != nil {
		return utils.WriteMessage(info.Conn, "425 Can't open data connection.\r\n")
	}

	info.DataConnection = conn
	return nil
}
