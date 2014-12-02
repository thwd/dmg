package dmg

type MappingParser struct {
	Parser  Parser
	Mapping func(interface{}) interface{}
}

func NewMappingParser(p Parser, m func(interface{}) interface{}) Parser {
	return MappingParser{p, m}
}

func (p MappingParser) Parse(bs []byte) StateSet {

	r := p.Parser.Parse(bs)

	return r.Map(func(s State) State {

		if s.Value == nil {
			return NewState(
				s.Remnant,
				nil,
				NewMappingParser(s.Parser, p.Mapping),
			)
		}

		if v, k := s.Value.(Accept); k {
			return NewState(
				s.Remnant,
				Accept{(p.Mapping)(v.Value)},
				s.Parser,
			)
		}

		return s
	})
}

func (p MappingParser) GoString() string {
	return "M(" + p.Parser.GoString() + ")"
}
