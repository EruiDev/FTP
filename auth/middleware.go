package auth

import (
	"myftp/commons"
	"myftp/utils"
)

func HandleSyst(args []string, user *commons.Info) error {
	return utils.WriteMessage(user.Conn, "215 UNIX Type: L8\r\n")
}

func HandleFeat(args []string, info *commons.Info) error {
	if err := utils.WriteMessage(info.Conn, "211-Features:\n"); err != nil {
		return err
	}
	return utils.WriteMessage(info.Conn, "211 End\n")
}
