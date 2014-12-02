package dmg

func NewMaybeParser(p Parser) Parser {
	return NewAlternationParser(
		p, NewEpsilonParser(),
	)
}
