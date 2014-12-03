package dmg

type MappingParser struct {
	Parser  Parser
	Mapping func(interface{}) interface{}
}

func NewMappingParser(p Parser, m func(interface{}) interface{}) Parser {
	return MappingParser{p, m}
}

func (p MappingParser) Parse(bs Remnant) StateSet {

	r := p.Parser.Parse(bs)

	return r.Map(func(s State) State {

		if !s.Final {
			return NewContinuedState(
				NewMappingParser(s.Parser, p.Mapping),
				s.Remnant,
			)
		}

		if s.Value.Success {
			return NewFinalState(
				Accept((p.Mapping)(s.Value.Value)),
				s.Remnant,
			)
		}

		return s
	})
}
