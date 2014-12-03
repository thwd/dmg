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

func NewPlusParser(p Parser) Parser {
	return NewSequenceParser(
		p, NewKleeneParser(p),
	)
}
