package dmg

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
