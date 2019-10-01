package commands

import (
	"errors"

	prompt "github.com/c-bata/go-prompt"
)

type GoQueryCommand func(string) error

// Function Map
var CommandMap map[string]GoQueryCommand
var SuggestionsMap []prompt.Suggest

// Errors
var errArgumentError error
var errRuntimeError error

func init() {
	CommandMap = map[string]GoQueryCommand{
		".connect":    connect,
		".disconnect": disconnect,
		".exit":       exit,
	}
	SuggestionsMap = []prompt.Suggest{
		{".connect", "Connect to a host with UUID"},
		{".disconnect", "Disconnect form a host with UUID"},
		{".exit", "Exit goquery"},
	}

	errArgumentError = errors.New("The arguments provided were incorrect for the command")
	errRuntimeError = errors.New("There was a problem executing the command")
}
