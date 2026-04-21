package commands

import (
	"fmt"
	"myftp/commons"
	"myftp/utils"
	"net"
	"strconv"
	"strings"
	"time"
)

func HandlePassive(args []string, info *commons.Info) error {
	if !info.IsLogged {
		return utils.WriteMessage(info.Conn, commons.LoginFirst)
	}

	if info.DataConnection != nil {
		_ = info.DataConnection.Close()
		info.DataConnection = nil
	}

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return utils.WriteMessage(info.Conn, commons.CantOpenDataConn)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	defer func() { _ = listener.Close() }()

	localAddr := info.Conn.LocalAddr().String()
	host, _, _ := net.SplitHostPort(localAddr)
	ipParts := strings.Split(host, ".")
	if len(ipParts) != 4 {
		ipParts = []string{"127", "0", "0", "1"}
	}

	pasvAddr := strings.Join(ipParts, ",") + "," + strconv.Itoa(port>>8) + "," + strconv.Itoa(port&0xff)
	if err := utils.WriteMessage(info.Conn, fmt.Sprintf(commons.EnteringPassiveMode, pasvAddr)); err != nil {
		return err
	}

	_ = listener.(*net.TCPListener).SetDeadline(time.Now().Add(30 * time.Second))
	conn, err := listener.Accept()
	if err != nil {
		return utils.WriteMessage(info.Conn, commons.CantOpenDataConn)
	}

	info.DataConnection = conn
	return nil
}
