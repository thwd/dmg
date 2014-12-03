package dmg

func NewMaybeParser(p Parser) Parser {
	return NewAlternationParser(
		p, NewEpsilonParser(),
	)
}

func NewKleeneParser(p Parser) Parser {
	return NewRecursiveParser(func(r Parser) Parser {
		return NewMaybeParser(
			NewSequenceParser(p, r),
		)
	})
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
