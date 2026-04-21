package auth

import (
	"fmt"
	"myftp/commons"
	"myftp/utils"
)

func HandleSyst(args []string, user *commons.Info) error {
	return utils.WriteMessage(user.Conn, commons.LinuxSystem)
}

func HandleFeat(args []string, info *commons.Info) error {
	if err := utils.WriteMessage(info.Conn, commons.FeatStart); err != nil {
		return err
	}
	return utils.WriteMessage(info.Conn, commons.FeatEnd)
}

func HandleType(args []string, info *commons.Info) error {
	if len(args) != 2 {
		return utils.WriteMessage(info.Conn, commons.InvalidCommand)
	}

	if args[1] == "A" || args[1] == "I" {
		return utils.WriteMessage(info.Conn, fmt.Sprintf(commons.TypeSet, args[1]))
	}
	return utils.WriteMessage(info.Conn, commons.TypeNotSupported)
}
