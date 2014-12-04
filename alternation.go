package dmg

// An AlternationParser is a Parser that will match a set of Parsers
// against a Remnant in no particular order and return the merged
// set of states again in no particular order. It will return
// a set of either all accepted or all rejected States.
//
// It panics if called with zero arguments.
type AlternationParser []Parser

func NewAlternationParser(ps ...Parser) Parser {

	if len(ps) == 0 {
		panic("NewAlternationParser called with zero parsers")
	}

	if len(ps) == 1 {
		return ps[0]
	}

	qs := make([]Parser, 0, len(ps))

	// merge alternations together
	for _, p := range ps {
		if a, k := p.(AlternationParser); k {
			for _, q := range a {
				qs = append(qs, q)
			}
		} else {
			qs = append(qs, p)
		}
	}

	return AlternationParser(qs)
}

func (p AlternationParser) Parse(bs Remnant) StateSet {

	passups, rejects := NewStateSet(), NewStateSet()

	for _, q := range p {

		r := q.Parse(bs)

		for r.Len() > 0 {

			s := r.Next()

			if s.Rejected() {
				rejects.Add(s)
			} else {
				passups.Add(s)
			}
		}
	}

	if passups.Len() == 0 {
		return rejects
	}

	return passups
}
