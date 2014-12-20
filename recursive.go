package dmg

// RecursiveParser is a Parser that allows the definition of recursive
// grammar trees. See NewRecursiveParser.
type RecursiveParser struct {
	Parser Parser
}

// NewRecursiveParser passes a pointer to a RecursiveParser to its
// first argument and defines said RecursiveParser as the result of
// this operation.
func NewRecursiveParser(f func(Parser) Parser) Parser {
	p := &RecursiveParser{}
	p.Parser = f(p)
	return p
}

// Parse delegates to the RecursiveParser's underlying parser.
func (p *RecursiveParser) Parse(bs Remnant) StateSet {
	return NewStateSet(
		Continue(p.Parser, bs),
	)
}

// NewMutuallyRecursiveParsers allows the definition of mutually recursive
// grammar trees.
//
// It creates an amount of RecursiveParsers equal to its first argument
// and passes them as a slice to its second argument. The result of this
// operation must be a slice of the same length (n), otherwise it panics.
//
// It returns a slice of length n of recursive Parsers with the order
// corresponding to the order returned by f.
func NewMutuallyRecursiveParsers(n int, f func([]Parser) []Parser) []Parser {

	ps := make([]Parser, n, n)

	for j := 0; j < n; j++ {
		ps[j] = &RecursiveParser{}
	}

	qs := f(ps)

	if len(qs) != len(ps) {
		panic("len(qs) != len(ps)")
	}

	for j := 0; j < n; j++ {
		ps[j].(*RecursiveParser).Parser = qs[j]
	}

	return ps

}
