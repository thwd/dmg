package dmg

type PrependParser struct {
	Prepend interface{}
	Parser  Parser
}

func NewPrependParser(r interface{}, p Parser) Parser {
	return PrependParser{r, p}
}

func (p PrependParser) Parse(bs Remnant) StateSet {

	r := p.Parser.Parse(bs)

	return r.Map(func(s State) State {

		if s.Continued() {
			return Continue(
				NewPrependParser(p.Prepend, s.Parser),
				s.Remnant,
			)
		}

		if s.Accepted() {
			return Accept(
				[2]interface{}{p.Prepend, s.Value},
				s.Remnant,
			)
		}

		return Reject(
			[2]interface{}{p.Prepend, s.Value},
			s.Remnant,
		)

	})
}
