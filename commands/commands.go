/*
Copyright © 2024 Justin Chen justin928501@gmail.com

*/
package commands

import (
	"github.com/spf13/cobra"
)

type commandsBuilder struct {
	commands []*cobra.Command
}

func newCommandsBuilder() *commandsBuilder {
	return &commandsBuilder{}
}

func (b *commandsBuilder) addCommands(commands ...*cobra.Command) *commandsBuilder {
	b.commands = append(b.commands, commands...)
	return b
}

func (b *commandsBuilder) addAll() *commandsBuilder {
	b.addCommands(
		b.newScrapCmd(),
	)

	return b
}

func (b *commandsBuilder) build() *bookwormCmd {
	bookworm := b.newBookwormCmd()

	for _, command := range b.commands {
		if command == nil {
			continue
		}
		bookworm.getCommand().AddCommand(command)
	}

	return bookworm
}

// Execute adds all child commands to the root command HugoCmd and sets flags appropriately.
// The args are usually filled with os.Args[1:].
func Execute(args []string) (*cobra.Command, error) {
	bookwormCmd := newCommandsBuilder().addAll().build()
	command := bookwormCmd.getCommand()
	command.SetArgs(args)
	return command.ExecuteC()
}
