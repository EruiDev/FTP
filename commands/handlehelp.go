package commands

import (
	"myftp/commons"
	"myftp/utils"
	"strings"
)

func HandleHelp(args []string, info *commons.Info, commandList map[string]commons.Command) error {
	var builder strings.Builder
	builder.WriteString("214-The following commands are recognized:\r\n")

	commands := make([]string, 0, len(commandList))
	for cmd := range commandList {
		commands = append(commands, cmd)
	}

	for _, cmd := range commands {
		desc := commandList[cmd].Description
		builder.WriteString("  ")
		builder.WriteString(cmd)
		if desc != "" {
			builder.WriteString(" - ")
			builder.WriteString(desc)
		}
		builder.WriteString("\r\n")
	}

	builder.WriteString("214 Help OK.\r\n")
	return utils.WriteMessage(info.Conn, builder.String())
}
