package commands

import (
	"io"
	"myftp/commons"
	"myftp/utils"
	"os"
	"path/filepath"
	"strings"
)

func HandleStor(args []string, info *commons.Info) error {
	if !info.IsLogged {
		return utils.WriteMessage(info.Conn, commons.LoginFirst)
	}
	if len(args) != 2 {
		return utils.WriteMessage(info.Conn, commons.InvalidCommand)
	}
	if info.DataConnection == nil {
		return utils.WriteMessage(info.Conn, commons.NoDataConnection)
	}

	defer func() {
		if info.DataConnection != nil {
			info.DataConnection.Close()
			info.DataConnection = nil
		}
	}()

	var targetPath string
	if filepath.IsAbs(args[1]) {
		targetPath = filepath.Clean(args[1])
	} else {
		targetPath = filepath.Clean(filepath.Join(info.CurrentDir, args[1]))
	}

	if !strings.HasPrefix(targetPath, info.OriginalDir+string(filepath.Separator)) && targetPath != info.OriginalDir {
		return utils.WriteMessage(info.Conn, commons.FileNotFound)
	}

	realDir, err := filepath.EvalSymlinks(filepath.Dir(targetPath))
	if err != nil {
		return utils.WriteMessage(info.Conn, commons.FileNotFound)
	}
	targetPath = filepath.Join(realDir, filepath.Base(targetPath))
	if !strings.HasPrefix(targetPath, info.OriginalDir+string(filepath.Separator)) && targetPath != info.OriginalDir {
		return utils.WriteMessage(info.Conn, commons.FileNotFound)
	}

	file, err := os.Create(targetPath)
	if err != nil {
		return utils.WriteMessage(info.Conn, commons.FileOpenError)
	}
	defer file.Close()

	err = utils.WriteMessage(info.Conn, commons.FileTransferStarting)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, info.DataConnection)

	if err != nil {
		return utils.WriteMessage(info.Conn, commons.FileTransferError)
	}

	return utils.WriteMessage(info.Conn, commons.FileTransferComplete)
}
