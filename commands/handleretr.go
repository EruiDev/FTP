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

	defer func() {
		if info.DataConnection != nil {
			_ = info.DataConnection.Close()
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

	realPath, err := filepath.EvalSymlinks(targetPath)
	if err != nil {
		return utils.WriteMessage(info.Conn, commons.FileNotFound)
	}
	if !strings.HasPrefix(realPath, info.OriginalDir+string(filepath.Separator)) && realPath != info.OriginalDir {
		return utils.WriteMessage(info.Conn, commons.FileNotFound)
	}

	fileInfo, err := os.Stat(realPath)
	if os.IsNotExist(err) || err != nil {
		return utils.WriteMessage(info.Conn, commons.FileNotFound)
	}
	if fileInfo.IsDir() {
		return utils.WriteMessage(info.Conn, commons.FileNotFound)
	}

	file, err := os.Open(realPath)
	if err != nil {
		return utils.WriteMessage(info.Conn, commons.FileOpenError)
	}
	defer func() { _ = file.Close() }()

	err = utils.WriteMessage(info.Conn, commons.FileTransferStarting)
	if err != nil {
		return err
	}

	_, err = io.Copy(info.DataConnection, file)

	if err != nil {
		return utils.WriteMessage(info.Conn, commons.FileTransferError)
	}

	return utils.WriteMessage(info.Conn, commons.FileTransferComplete)
}
