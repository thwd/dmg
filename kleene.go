package dmg

func NewKleeneParser(p Parser) Parser {
	return NewRecursiveParser(func(r Parser) Parser {
		return NewMaybeParser(
			NewSequenceParser(p, r),
		)
	})
}
