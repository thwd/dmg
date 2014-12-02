package dmg

type SequenceParser []Parser

func NewSequenceParser(ps ...Parser) Parser {
	if len(ps) == 0 {
		panic("NewSequenceParser called with zero parsers")
	}
	if len(ps) == 1 {
		return ps[0]
	}
	return SequenceParser(ps)
}

func (p SequenceParser) Parse(bs []byte) StateSet {

	passups, rejects := NewStateSet(), NewStateSet()

	r := p[0].Parse(bs)

	for i, l := 0, r.Len(); i < l; i++ {

		s := r.Next()

		if s.Value == nil {

			cont := append([]Parser{s.Parser}, p[1:]...)

			passups.Add(
				NewState(
					s.Remnant,
					nil,
					NewSequenceParser(cont...),
				),
			)
			continue
		}

		if v, k := s.Value.(Accept); k {
			passups.Add(
				NewState(
					s.Remnant,
					nil,
					NewPrependParser(v.Value, NewSequenceParser(p[1:]...)),
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

func (p SequenceParser) GoString() string {
	t := ""
	for _, q := range p {
		t += q.GoString() + "·"
	}
	t = t[:len(t)-2] // len("·") == 2
	return t
}
