package dmg

import (
	"fmt"
)

type State struct {
	state   byte
	Remnant Remnant
	Parser  Parser
	Value   interface{}
}

const (
	stateAccepted byte = iota
	stateRejected
	stateContinued
)

func Accept(v interface{}, r Remnant) State {
	return State{stateAccepted, r, nil, v}
}

func Reject(v interface{}, r Remnant) State {
	return State{stateRejected, r, nil, v}
}

func Continue(p Parser, r Remnant) State {
	return State{stateContinued, r, p, nil}
}

func (s State) Reduce() StateSet {
	return s.Parser.Parse(s.Remnant)
}

func (s State) Continued() bool {
	return (s.state == stateContinued)
}

func (s State) Accepted() bool {
	return (s.state == stateAccepted)
}

func (s State) Rejected() bool {
	return (s.state == stateRejected)
}

func (s State) GoString() string {
	t := "{\n"
	t += "\tRemnant: \"" + string(s.Remnant) + "\"\n"
	t += "\tParser:  " + fmt.Sprintf("%#v", s.Parser) + "\n"
	t += "\tValue:   " + fmt.Sprintf("%#v", s.Value) + "\n"
	t += "}\n"
	return t
}
