package dmg

// A SequenceParser is a Parser that will match a set of Parsers
// against a Remnant one after the other.
type SequenceParser []Parser

// NewSequenceParser returns a new SequenceParser.
// It panics if called with zero arguments.
func NewSequenceParser(ps ...Parser) Parser {
	if len(ps) == 0 {
		panic("NewSequenceParser called with zero parsers")
	}
	if len(ps) == 1 {
		return ps[0]
	}
	return SequenceParser(ps)
}

// Parse applies every element in a sequence to a remnant, in order.
func (p SequenceParser) Parse(r Remnant, c chan State) {

	d, f := make(chan State), make(chan struct{})

	go func() {
		for s := range d {

			if s.Continued() {
				c <- Continue(
					NewSequenceParser(
						append([]Parser{s.Parser}, p[1:]...)...,
					),
					s.Remnant,
				)
				continue
			}

			if s.Accepted() {
				c <- Continue(
					NewPrependParser(s.Value, NewSequenceParser(p[1:]...)),
					s.Remnant,
				)
				continue
			}

			c <- s

		}
		f <- struct{}{}
	}()

	p[0].Parse(r, d)

	close(d)

	<-f
}
