package commands

import (
	"myftp/commons"
	"myftp/utils"
	"net"
	"os"
	"path/filepath"
)

func displayFiles(files []os.DirEntry, conn net.Conn) error {
	for _, file := range files {
		err := utils.WriteMessage(conn, file.Name()+"\r\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func HandleList(args []string, info *commons.Info) error {
	if !info.IsLogged {
		return utils.WriteMessage(info.Conn, commons.LoginFirst)
	}
	if len(args) > 2 {
		return utils.WriteMessage(info.Conn, commons.InvalidCommand)
	}
	if info.DataConnection == nil {
		return utils.WriteMessage(info.Conn, commons.NoDataConnection)
	}

	path := "."
	if len(args) == 2 {
		path = args[1]
	}

	absPath, err := filepath.Abs(filepath.Join(info.CurrentDir, path))
	if err != nil {
		return utils.WriteMessage(info.Conn, "550 Invalid path\r\n")
	}

	files, err := os.ReadDir(absPath)
	if err != nil {
		return utils.WriteMessage(info.Conn, commons.DirectoryNotFound)
	}

	err = utils.WriteMessage(info.Conn, "150 Here comes the directory listing.\r\n")
	if err != nil {
		return utils.WriteMessage(info.Conn, "550 Failed to list directory\r\n")
	}

	err = displayFiles(files, info.DataConnection)
	info.DataConnection.Close()
	info.DataConnection = nil
	if err != nil {
		return utils.WriteMessage(info.Conn, "550 Failed to list all directory")
	}
	err = utils.WriteMessage(info.Conn, "226 Directory send OK.\r\n")
	if err != nil {
		return err
	}
	return nil
}
