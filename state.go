package dmg

// A State represents a step in the parsing process. It can be either
// accepted, rejected or continued, meaning that either it matched an
// input, didn't match an input or has yet to match an input,
// respectively.
//
// If a State is accepted or rejected, its Parser field will be nil and
// its Value field not nil. The opposite otherwise.
//
// A State can be accepted or rejected even if it has a non-empty Remnant.
type State struct {
	state   byte
	Remnant Remnant
	Parser  Parser
	Value   interface{}
}

const (
	stateContinued byte = iota
	stateAccepted
	stateRejected
)

// Accept returns a accepted state with Value v and Remnant r
func Accept(v interface{}, r Remnant) State {
	return State{stateAccepted, r, nil, v}
}

// Reject returns a rejected state with Value v and Remnant r
func Reject(v interface{}, r Remnant) State {
	return State{stateRejected, r, nil, v}
}

// Continue returns a continued state with Parser p and Remnant r
func Continue(p Parser, r Remnant) State {
	return State{stateContinued, r, p, nil}
}

// Reduce applies a State's Parser to it's Remnant and returns the
// resulting StateSet.
func (s State) Reduce() StateSet {
	return s.Parser.Parse(s.Remnant)
}

// Continued reports wether a State is continued.
func (s State) Continued() bool {
	return (s.state == stateContinued)
}

// Accepted reports wether a State is accepted.
func (s State) Accepted() bool {
	return (s.state == stateAccepted)
}

// Rejected reports wether a State is rejected.
func (s State) Rejected() bool {
	return (s.state == stateRejected)
}
