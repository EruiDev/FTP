package commands

import (
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
	return utils.WriteMessage(info.Conn, "257 \""+info.CurrentDir+"\" is the current directory \r\n")
}
