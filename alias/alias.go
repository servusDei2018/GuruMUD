package alias

import (
	"errors"
	"strings"
)

// An Alias represents an alias.
type Alias struct {
	// This Alias's name.
	Name string
	// The command this Alias represents.
	Command string
}

// NewAlias returns a new Alias.
func NewAlias(name, cmd string) (*Alias, error) {
	if len(name) < 3 {
		return &Alias{}, errors.New("alias: name is too short")
	}
	if len(cmd) < 1 {
		return &Alias{}, errors.New("alias: command is too short")
	}

	return &Alias{Name: name, Command: cmd}, nil
}

// Execute return a list of the alias' commands.
func (a *Alias) Execute() []string {
	return strings.Split(a.Command, ";")
}
