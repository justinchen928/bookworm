package commands

import "github.com/spf13/cobra"

func (cc *bookwormCmd) getCommand() *cobra.Command {
	return cc.command
}

type bookwormCmd struct {
	command *cobra.Command
}

func (b *commandsBuilder) newBookwormCmd() *bookwormCmd {

	bookwormCmd := &bookwormCmd{}

	cmd := &cobra.Command{
		Use: "bookworm",
		Short: `Bookworm is a tool to scrape web novels of certain websites and save to file.
Also a tool to transform novel in txt to json format to e-book file format like epub or kepub.`,
	}

	cmd.CompletionOptions.DisableDefaultCmd = true

	bookwormCmd.command = cmd

	return bookwormCmd
}
