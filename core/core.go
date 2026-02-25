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
	"HELP": {
		Description: "Display all commands available with a description",
		Handler:     nil, // Handled specially in protocol.go
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
	"LIST": {
		Description: "List the files in the current directory",
		Handler:     commands.HandleList,
	},
	"CWD": {
		Description: "Change the current directory",
		Handler:     commands.HandleCwd,
	},
	"CDUP": {
		Description: "Change to parent directory",
		Handler:     commands.HandleCdup,
	},
	"PASV": {
		Description: "Enter passive mode",
		Handler:     commands.HandlePassive,
	},
	"PORT": {
		Description: "Specify the client data port",
		Handler:     commands.HandlePort,
	},
	"DELE": {
		Description: "Delete a file",
		Handler:     commands.HandleDele,
	},
	"RETR": {
		Description: "Retrieve (download) a file",
		Handler:     commands.HandleRetr,
	},
	"STOR": {
		Description: "Store (upload) a file",
		Handler:     commands.HandleStor,
	},
	"NOOP": {
		Description: "No operation (do nothing)",
		Handler:     commands.HandleNoop,
	},
	"TYPE": {
		Description: "Provide the file type to the client",
		Handler:     auth.HandleType,
	},
}
