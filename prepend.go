package dmg

type PrependParser struct {
	Prepend interface{}
	Parser  Parser
}

func NewPrependParser(r interface{}, p Parser) Parser {
	return PrependParser{r, p}
}

func (p PrependParser) Parse(bs []byte) StateSet {

	r := p.Parser.Parse(bs)

	return r.Map(func(s State) State {

		if s.Value == nil {
			return NewState(
				s.Remnant,
				nil,
				NewPrependParser(p.Prepend, s.Parser),
			)
		}

		if v, k := s.Value.(Accept); k {
			return NewState(
				s.Remnant,
				Accept{[2]interface{}{p.Prepend, v.Value}},
				s.Parser,
			)
		}

		return NewState(
			s.Remnant,
			Reject{[2]interface{}{p.Prepend, s.Value.(Reject).Value}},
			s.Parser,
		)

	})
}

func (p PrependParser) GoString() string {
	return "P(" + p.Parser.GoString() + ")"
}
