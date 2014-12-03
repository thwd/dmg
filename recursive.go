package dmg

type RecursiveParser struct {
	Parser Parser
}

func NewRecursiveParser(f func(Parser) Parser) Parser {
	p := &RecursiveParser{}
	p.Parser = f(p)
	return p
}

func (p *RecursiveParser) Parse(bs Remnant) StateSet {
	return NewStateSet(
		NewContinuedState(p.Parser, bs),
	)
}
