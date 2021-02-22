package trigger

import (
	"errors"
	"strings"
)

type Trigger struct {
	Name string
	Trigger string
	Response string
}

// NewTrigger returns a new trigger.
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
