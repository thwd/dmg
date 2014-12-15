package dmg

// NewMaybeParser returns a Parser that accepts either whatever
// its first argument accepts or the empty string.
func NewMaybeParser(p Parser) Parser {
	return NewAlternationParser(
		p, NewEpsilonParser(),
	)
}

// NewKleeneParser returns a Parser that accepts whatever
// its first argument accepts zero or many times.
func NewKleeneParser(p Parser) Parser {
	return NewRecursiveParser(func(r Parser) Parser {
		return NewMaybeParser(
			NewSequenceParser(p, r),
		)
	})
}

// NewPlusParser returns a Parser that accepts whatever
// its first argument accepts one or many times.
func NewPlusParser(p Parser) Parser {
	return NewSequenceParser(
		p, NewKleeneParser(p),
	)
}

func MatchToString(m interface{}) string {
	if m == nil {
		return ""
	}
	if v, k := m.([2]interface{}); k {
		return MatchToString(v[0]) + MatchToString(v[1])
	}
	return string(m.(Remnant))
}
