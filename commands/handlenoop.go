package commands

import (
	"myftp/commons"
	"myftp/utils"
)

func HandleNoop(args []string, info *commons.Info) error {
	return utils.WriteMessage(info.Conn, "200 OK.\r\n")
}
