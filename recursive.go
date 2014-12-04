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
		Continue(p.Parser, bs),
	)
}

func NewMutuallyRecursiveParsers(i int, f func([]Parser) []Parser) []Parser {

	ps := make([]Parser, i, i)

	for j := 0; j < i; j++ {
		ps[j] = &RecursiveParser{}
	}

	qs := f(ps)

	if len(qs) != len(ps) {
		panic("len(qs) != len(ps)")
	}

	for j := 0; j < i; j++ {
		ps[j].(*RecursiveParser).Parser = qs[j]
	}

	return ps

}
