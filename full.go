package dmg

type FullParser struct {
	Parser Parser
}

func NewFullParser(p Parser) Parser {
	return FullParser{p}
}

func (p FullParser) GoString() string {
	return "F(" + p.GoString() + ")"
}

func (p FullParser) Parse(bs []byte) StateSet {

	q := NewStateSet(NewState(bs, nil, p.Parser))

	for {

		r := q.Next().Reduce()

		for i, l := 0, r.Len(); i < l; i++ {

			t := r.Next()

			if t.Parser == nil && len(t.Remnant) == 0 {
				return NewStateSet(t)
			}

			if t.Parser == nil {
				continue
			}

			q.Add(t)

		}

	}
}
