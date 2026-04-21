package commands

import (
	"fmt"
	"myftp/commons"
	"myftp/utils"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func formatDirEntry(entry os.DirEntry) (string, error) {
	info, err := entry.Info()
	if err != nil {
		return "", err
	}
	modTime := info.ModTime().Format("Jan 02 15:04")
	return fmt.Sprintf("%s 1 ftp ftp %12d %s %s\r\n",
		info.Mode().String(), info.Size(), modTime, entry.Name()), nil
}

func displayFiles(files []os.DirEntry, conn net.Conn) error {
	for _, file := range files {
		line, err := formatDirEntry(file)
		if err != nil {
			line = file.Name() + "\r\n"
		}
		if err := utils.WriteMessage(conn, line); err != nil {
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

	defer func() {
		if info.DataConnection != nil {
			info.DataConnection.Close()
			info.DataConnection = nil
		}
	}()

	path := "."
	if len(args) == 2 {
		path = args[1]
	}

	absPath, err := filepath.Abs(filepath.Join(info.CurrentDir, path))
	if err != nil {
		return utils.WriteMessage(info.Conn, commons.InvalidPath)
	}

	if !strings.HasPrefix(absPath, info.OriginalDir+string(filepath.Separator)) && absPath != info.OriginalDir {
		return utils.WriteMessage(info.Conn, commons.DirectoryNotFound)
	}

	files, err := os.ReadDir(absPath)
	if err != nil {
		return utils.WriteMessage(info.Conn, commons.DirectoryNotFound)
	}

	if err := utils.WriteMessage(info.Conn, commons.DirectoryListing); err != nil {
		return err
	}

	if err := displayFiles(files, info.DataConnection); err != nil {
		return utils.WriteMessage(info.Conn, commons.ListFailed)
	}

	return utils.WriteMessage(info.Conn, commons.DirectoryOK)
}
