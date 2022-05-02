package commands

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNewCommandsBuilder(t *testing.T) {
	assert.Equal(t, &commandsBuilder{}, newCommandsBuilder())
}

func TestAddCommands(t *testing.T) {
	commandsBuilder := &commandsBuilder{}
	commandsBuilder.addCommands(&cobra.Command{}, &cobra.Command{})
	assert.Equal(t, 2, len(commandsBuilder.commands))
}

func TestAddAll(t *testing.T) {
	commandsBuilder := &commandsBuilder{}
	commandsBuilder.addAll()
	assert.Equal(t, 1, len(commandsBuilder.commands))
}

func TestBuildNoSubCommands(t *testing.T) {
	commandsBuilder := &commandsBuilder{}
	bookwormCmd := commandsBuilder.build()
	assert.Equal(t, false, bookwormCmd.getCommand().HasSubCommands())
}

func TestBuildHasSubCommands(t *testing.T) {
	commandsBuilder := &commandsBuilder{}
	bookwormCmd := commandsBuilder.addAll().build()
	assert.Equal(t, true, bookwormCmd.getCommand().HasSubCommands())
}
