package dmg

import (
	"fmt"
	"runtime"
)

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
	status  byte
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

func worker(in chan State, out chan State) {
	for {
		s := <-in
		s.Parser.Parse(s.Remnant, out)
	}
}

func tee(in, cont chan State, done chan State) {
	for {
		p := <-in

		if p.Continued() {
			cont <- p
			continue
		}

		done <- p
	}
}

func printQueue(q []State) {
	for _, e := range q {
		fmt.Printf("%#v,", len(e.Remnant))
	}
	fmt.Println("")
}

func order(in chan State, out chan State) {

	q := make([]State, 0)

	for {

	read_loop:
		for {
			select {
			case p := <-in:

				q = append(q, p)

				for i := (len(q) - 1); i > 0; i-- {

					h, l := q[i], q[i-1]

					if len(l.Remnant) > len(h.Remnant) {
						break
					}

					q[i], q[i-1] = l, h
				}

			default:
				break read_loop
			}
		}

	write_loop:
		for len(q) > 0 {
			select {
			case out <- q[len(q)-1]:
				q = q[:len(q)-1]
			default:
				break write_loop
			}
		}

		runtime.Gosched()
	}

}

// Reduce applies a State's Parser to its Remnant and returns the
// resulting states in a channel.
func (s State) Reduce(c chan State, w int) {

	d, e, f := make(chan State), make(chan State), make(chan State)

	go tee(d, e, c)

	go order(e, f)

	for i := 0; i < w; i++ {
		go worker(f, d)
	}

	f <- s
}

// Continued reports wether a State is continued.
func (s State) Continued() bool {
	return (s.status == stateContinued)
}

// Accepted reports wether a State is accepted.
func (s State) Accepted() bool {
	return (s.status == stateAccepted)
}

// Rejected reports wether a State is rejected.
func (s State) Rejected() bool {
	return (s.status == stateRejected)
}
