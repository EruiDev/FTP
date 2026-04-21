package commands

import (
	"myftp/commons"
	"myftp/utils"
	"os"
	"path/filepath"
	"strings"
)

func HandleCwd(args []string, info *commons.Info) error {
	if !info.IsLogged {
		return utils.WriteMessage(info.Conn, commons.LoginFirst)
	}
	if len(args) != 2 {
		return utils.WriteMessage(info.Conn, commons.InvalidCommand)
	}

	var newDir string
	if filepath.IsAbs(args[1]) {
		newDir = filepath.Clean(args[1])
	} else {
		newDir = filepath.Clean(filepath.Join(info.CurrentDir, args[1]))
	}

	if !strings.HasPrefix(newDir, info.OriginalDir+string(filepath.Separator)) && newDir != info.OriginalDir {
		return utils.WriteMessage(info.Conn, commons.DirectoryNotFound)
	}

	realDir, err := filepath.EvalSymlinks(newDir)
	if err != nil {
		return utils.WriteMessage(info.Conn, commons.DirectoryNotFound)
	}
	if !strings.HasPrefix(realDir, info.OriginalDir+string(filepath.Separator)) && realDir != info.OriginalDir {
		return utils.WriteMessage(info.Conn, commons.DirectoryNotFound)
	}

	fileInfo, err := os.Stat(realDir)
	if os.IsNotExist(err) || err != nil || !fileInfo.IsDir() {
		return utils.WriteMessage(info.Conn, commons.DirectoryNotFound)
	}

	info.CurrentDir = realDir
	return utils.WriteMessage(info.Conn, commons.DirectoryChanged)
}

func HandleCdup(args []string, info *commons.Info) error {
	if !info.IsLogged {
		return utils.WriteMessage(info.Conn, commons.LoginFirst)
	}

	parentDir := filepath.Dir(info.CurrentDir)

	if !strings.HasPrefix(parentDir, info.OriginalDir+string(filepath.Separator)) && parentDir != info.OriginalDir {
		return utils.WriteMessage(info.Conn, commons.DirectoryNotFound)
	}

	info.CurrentDir = parentDir
	return utils.WriteMessage(info.Conn, commons.DirectoryChanged)
}
