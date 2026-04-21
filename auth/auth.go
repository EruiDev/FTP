package auth

import (
	"myftp/commons"
	"myftp/utils"
	sys_usr "os/user"
	"time"
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

func HandlePass(args []string, info *commons.Info) error {
	if info.IsLogged {
		return utils.WriteMessage(info.Conn, commons.AlreadyLogged)
	}
	if info.Username == "" {
		return utils.WriteMessage(info.Conn, commons.UserFirst)
	}

	time.Sleep(time.Second)

	user, err := sys_usr.Lookup(info.Username)
	if err != nil {
		return utils.WriteMessage(info.Conn, commons.AuthenticationError)
	}

	ok, err := Authenticate(info.Username, args[1])
	if err != nil {
		return utils.WriteMessage(info.Conn, commons.AuthenticationError)
	}
	if !ok {
		return utils.WriteMessage(info.Conn, commons.IncorrectLogin)
	}
	info.OriginalDir = user.HomeDir
	info.CurrentDir = user.HomeDir
	info.IsLogged = true

	return utils.WriteMessage(info.Conn, commons.UserLoginSuccess)
}
