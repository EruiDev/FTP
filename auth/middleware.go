package auth

import (
	"myftp/commons"
	"myftp/utils"
)

func HandleSyst(args []string, user *commons.Info) error {
	return utils.WriteMessage(user.Conn, "215 UNIX Type: L8\r\n")
}

func HandleFeat(args []string, info *commons.Info) error {
	if err := utils.WriteMessage(info.Conn, "211-Features:\r\n"); err != nil {
		return err
	}
	return utils.WriteMessage(info.Conn, "211 End\r\n")
}

func HandleType(args []string, info *commons.Info) error {
	if len(args) != 2 {
		return utils.WriteMessage(info.Conn, "500 Invalid command.\r\n")
	}

	if args[1] == "A" || args[1] == "I" {
		return utils.WriteMessage(info.Conn, "200 Type set to "+args[1]+".\r\n")
	} else {
		return utils.WriteMessage(info.Conn, "504 Type not supported.\r\n")
	}
}
