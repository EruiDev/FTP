package utils

import "net"

func WriteMessage(conn net.Conn, message string) error {
	if _, err := conn.Write([]byte(message)); err != nil {
		return err
	}
	return nil
}
