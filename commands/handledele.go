package commands

import (
	"myftp/commons"
	"myftp/utils"
	"os"
	"path/filepath"
	"strings"
)

func HandleDele(args []string, info *commons.Info) error {
	if !info.IsLogged {
		return utils.WriteMessage(info.Conn, commons.LoginFirst)
	}
	if len(args) != 2 {
		return utils.WriteMessage(info.Conn, commons.InvalidCommand)
	}

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
	if os.IsNotExist(err) {
		return utils.WriteMessage(info.Conn, commons.FileNotFound)
	}
	if err != nil {
		return utils.WriteMessage(info.Conn, commons.FileNotFound)
	}

	if fileInfo.IsDir() {
		return utils.WriteMessage(info.Conn, commons.FileNotFound)
	}

	if err := os.Remove(realPath); err != nil {
		return utils.WriteMessage(info.Conn, commons.FileDeleteError)
	}

	return utils.WriteMessage(info.Conn, commons.FileDeleted)
}
