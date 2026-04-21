package commands

import (
	"fmt"
	"myftp/commons"
	"myftp/utils"
)

func HandlePwd(args []string, info *commons.Info) error {
	if !info.IsLogged {
		return utils.WriteMessage(info.Conn, commons.LoginFirst)
	}
	if info.CurrentDir == "" {
		return utils.WriteMessage(info.Conn, commons.NoDirectorySelected)
	}
	return utils.WriteMessage(info.Conn, fmt.Sprintf(commons.CurrentDirectory, info.CurrentDir))
}
