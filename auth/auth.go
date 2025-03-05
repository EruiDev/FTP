package auth

import (
	"myftp/commons"
	"myftp/utils"
)

func HandleUser(args []string, info *commons.Info) error {
	if info.IsLogged {
		return utils.WriteMessage(info.Conn, commons.AlreadyLogged)
	}
	if len(args) >= 2 {
		info.Username = args[1]
		if info.Username == "anonymous" {
			info.IsLogged = true
			return utils.WriteMessage(info.Conn, commons.UserLoginSuccess)
		}
		return utils.WriteMessage(info.Conn, commons.UsernameOK)
	} else {
		if err := utils.WriteMessage(info.Conn, commons.IncorrectLogin); err != nil {
			return err
		}
	}
	return nil
}

func HandlePass(args []string, user *commons.Info) error {
	if user.IsLogged {
		return utils.WriteMessage(user.Conn, commons.AlreadyLogged)
	}
	if user.Username == "" {
		return utils.WriteMessage(user.Conn, commons.UserFirst)
	}
	user.Username = ""
	return utils.WriteMessage(user.Conn, commons.IncorrectLogin)
	// TODO: implement login system with crypt()
	// if password incorrect 530 and clear PASS
}
