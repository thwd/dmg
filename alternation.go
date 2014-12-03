package dmg

type AlternationParser []Parser

func NewAlternationParser(ps ...Parser) Parser {
	if len(ps) == 0 {
		panic("NewAlternationParser called with zero parsers")
	}
	if len(ps) == 1 {
		return ps[0]
	}
	return AlternationParser(ps)
}

func (p AlternationParser) Parse(bs Remnant) StateSet {

	passups, rejects := NewStateSet(), NewStateSet()

	for _, q := range p {

		r := q.Parse(bs)

		for r.Len() > 0 {

			s := r.Next()

			if s.Final && !s.Value.Success {
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
