package core

import (
	"myftp/auth"
	"myftp/commands"
	"myftp/commons"
)

var commandList = map[string]commons.Command{
	"USER": {
		Description: "Specify the user for authentication",
		Handler:     auth.HandleUser,
	},
	"PASS": {
		Description: "Specify the password for authentication",
		Handler:     auth.HandlePass,
	},
	"LIST": {
		Description: "List files in the current directory",
		Handler:     commands.HandleList,
	},
	"HELP": {
		Description: "Display all commands available with a description",
		Handler:     nil,
	},
	"SYST": {
		Description: "Display the system type",
		Handler:     auth.HandleSyst,
	},
	"FEAT": {
		Description: "Display the features available",
		Handler:     auth.HandleFeat,
	},
	"PWD": {
		Description: "Display the current directory",
		Handler:     commands.HandlePwd,
	},
	"CWD": {
		Description: "Change the current directory",
		Handler:     commands.HandleCwd,
	},
	"PASV": {
		Description: "Enter passive mode",
		Handler:     commands.HandlePassive,
	},
}