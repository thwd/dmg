package dmg

// A SequenceParser is a Parser that will match a set of Parsers
// against a Remnant one after the other and return a set of either
// all accepted or all rejected States.
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

func (p SequenceParser) Parse(bs Remnant) StateSet {

	passups, rejects := NewStateSet(), NewStateSet()

	r := p[0].Parse(bs)

	for r.Len() > 0 {

		s := r.Next()

		if s.Continued() {

			cont := append([]Parser{s.Parser}, p[1:]...)

			passups.Add(
				Continue(
					NewSequenceParser(cont...),
					s.Remnant,
				),
			)

			continue
		}

		if s.Accepted() {

			passups.Add(
				Continue(
					NewPrependParser(s.Value, NewSequenceParser(p[1:]...)),
					s.Remnant,
				),
			)

			continue
		}

		rejects.Add(s)

	}

	if passups.Len() == 0 {
		return rejects
	}

	return passups
}
