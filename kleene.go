package dmg

type KleeneParser struct {
	Parser Parser
}

func NewKleeneParser(p Parser) Parser {
	return KleeneParser{p}
}

func (p KleeneParser) Parse(bs []byte) StateSet {

	r := NewState(bs, nil, p.Parser).Reduce()

	return r.Map(func(s State) State {

		if s.Value == nil {
			return NewState(
				s.Remnant,
				nil,
				NewMaybeParser(
					NewSequenceParser(
						s.Parser,
						p,
					),
				),
			)
		}

		if v, k := s.Value.(Accept); k {
			return NewState(
				s.Remnant,
				nil,
				NewPrependParser(v.Value, p),
			)
		}

		return NewState(
			bs,
			Accept{bs[:0]},
			nil,
		)
	})
}

func (p KleeneParser) GoString() string {
	return p.Parser.GoString() + "*"
}
