package dmg

// MappingParser is a Parser that matches a given Parser against a Remnant
// and mutates the Values of all accepted States through a mapping function.
type MappingParser struct {
	Parser  Parser
	Mapping func(interface{}) interface{}
}

func NewMappingParser(p Parser, m func(interface{}) interface{}) Parser {
	return MappingParser{p, m}
}

// Parse delegates to the MappingParser's underlying parser and maps all
// matched states' values through the MappingParser's mapping function.
func (p MappingParser) Parse(bs Remnant) StateSet {

	r := p.Parser.Parse(bs)

	return r.Map(func(s State) State {

		if s.Continued() {
			return Continue(
				NewMappingParser(s.Parser, p.Mapping),
				s.Remnant,
			)
		}

		if s.Accepted() {
			return Accept(
				(p.Mapping)(s.Value),
				s.Remnant,
			)
		}

		return s
	})
}
