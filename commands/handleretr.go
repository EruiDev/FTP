package commands

import (
	"io"
	"myftp/commons"
	"myftp/utils"
	"os"
	"path/filepath"
	"strings"
)

func HandleRetr(args []string, info *commons.Info) error {
	if !info.IsLogged {
		return utils.WriteMessage(info.Conn, commons.LoginFirst)
	}
	if len(args) != 2 {
		return utils.WriteMessage(info.Conn, commons.InvalidCommand)
	}
	if info.DataConnection == nil {
		return utils.WriteMessage(info.Conn, commons.NoDataConnection)
	}

	var targetPath string
	if filepath.IsAbs(args[1]) {
		targetPath = filepath.Clean(args[1])
	} else {
		targetPath = filepath.Clean(filepath.Join(info.CurrentDir, args[1]))
	}

	if !strings.HasPrefix(targetPath, info.OriginalDir+string(filepath.Separator)) && targetPath != info.OriginalDir {
		info.DataConnection.Close()
		info.DataConnection = nil
		return utils.WriteMessage(info.Conn, commons.FileNotFound)
	}

	fileInfo, err := os.Stat(targetPath)
	if os.IsNotExist(err) || err != nil {
		info.DataConnection.Close()
		info.DataConnection = nil
		return utils.WriteMessage(info.Conn, commons.FileNotFound)
	}
	if fileInfo.IsDir() {
		info.DataConnection.Close()
		info.DataConnection = nil
		return utils.WriteMessage(info.Conn, commons.FileNotFound)
	}

	file, err := os.Open(targetPath)
	if err != nil {
		info.DataConnection.Close()
		info.DataConnection = nil
		return utils.WriteMessage(info.Conn, commons.FileOpenError)
	}
	defer file.Close()

	err = utils.WriteMessage(info.Conn, commons.FileTransferStarting)
	if err != nil {
		info.DataConnection.Close()
		info.DataConnection = nil
		return err
	}

	_, err = io.Copy(info.DataConnection, file)
	info.DataConnection.Close()
	info.DataConnection = nil

	if err != nil {
		return utils.WriteMessage(info.Conn, commons.FileTransferError)
	}

	return utils.WriteMessage(info.Conn, commons.FileTransferComplete)
}
