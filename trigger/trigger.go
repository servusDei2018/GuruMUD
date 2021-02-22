package trigger

import (
	"errors"
	"strings"
)

// A Trigger represents a trigger.
type Trigger struct {
	// This Trigger's name.
	Name string
	// Text which activates this Trigger.
	Trigger string
	// A `;` delimited string containing response commands.
	Response string
}

// NewTrigger returns a new Trigger.
func NewTrigger(name, trigger, response string) (*Trigger, error) {
	if len(name)<3{return Trigger{}, errors.New("trigger: name too short")}
	if len(trigger)<3{return Trigger{}, errors.New("trigger: trigger too short")}
	if len(response)<1{return Trigger{}, errors.New("trigger: response too short")}

	return &Trigger{Name: name, Trigger: trigger, Response: response}, nil
}

// Check checks whether the trigger is activated.
func (t *Trigger) Check(s string) bool {
	if strings.Contains(s, t.Trigger) {return true}

	return false
}

// Execute returns the response(s) as an array.
func (t *Trigger) Execute() []string {
	return strings.Split(t.Response, ";")
}
