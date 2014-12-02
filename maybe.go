package dmg

type MaybeParser struct {
	Parser Parser
}

func NewMaybeParser(p Parser) Parser {
	return MaybeParser{p}
}

func (p MaybeParser) Parse(bs []byte) StateSet {

	r := NewState(bs, nil, p.Parser).Reduce()

	return r.Map(func(s State) State {

		if s.Value == nil {
			return NewState(
				s.Remnant,
				nil,
				NewMaybeParser(s.Parser),
			)
		}

		if _, k := s.Value.(Accept); k {
			return s
		}

		return NewState(
			bs,
			Accept{bs[:0]},
			nil,
		)
	})
}

func (p MaybeParser) GoString() string {
	return p.Parser.GoString() + "?"
}
