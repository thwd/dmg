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

func (p AlternationParser) Parse(bs []byte) StateSet {

	passups, rejects := NewStateSet(), NewStateSet()

	for _, q := range p {

		r := q.Parse(bs)

		for i, l := 0, r.Len(); i < l; i++ {

			s := r.Next()

			if _, k := s.Value.(Reject); k {
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

func (p AlternationParser) GoString() string {
	t := "("
	for _, q := range p {
		t += q.GoString() + "|"
	}
	t = t[:len(t)-1]
	return t + ")"
}
