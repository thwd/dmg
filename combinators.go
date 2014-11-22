package dmg

type AlternationParser []Parser

func NewAlternationParser(ps ...Parser) Parser {
	if len(ps) == 0 {
		panic("NewAlternationParser called with zero parsers")
	}
	if len(ps) == 1 {
		return ps[0]
	}
	return AlternationParser(ps)
}

func (p AlternationParser) Parse(bs []byte) StateSet {

	passups, rejects := NewStateSet(), NewStateSet()

	for _, q := range p {

		r := q.Parse(bs)

		for i, l := 0, r.Len(); i < l; i++ {

			s := r.Next()

			if _, k := s.Value.(Reject); k {
				rejects.Add(s)
			} else {
				passups.Add(s)
			}

		}
	}

	if passups.Len() == 0 {
		return rejects
	}

	return passups
}

func (p AlternationParser) GoString() string {
	t := "("
	for _, q := range p {
		t += q.GoString() + "|"
	}
	t = t[:len(t)-1]
	return t + ")"
}

type PrependParser struct {
	Prepend interface{}
	Parser  Parser
}

func NewPrependParser(r interface{}, p Parser) Parser {
	return PrependParser{r, p}
}

func (p PrependParser) Parse(bs []byte) StateSet {

	r := p.Parser.Parse(bs)

	return r.Map(func(s State) State {

		if s.Value == nil {
			return NewState(
				s.Remnant,
				nil,
				NewPrependParser(p.Prepend, s.Parser),
			)
		}

		if v, k := s.Value.(Accept); k {
			return NewState(
				s.Remnant,
				Accept{[2]interface{}{p.Prepend, v.Value}},
				s.Parser,
			)
		}

		return NewState(
			s.Remnant,
			Reject{[2]interface{}{p.Prepend, s.Value.(Reject).Value}},
			s.Parser,
		)

	})
}

func (p PrependParser) GoString() string {
	return "P(" + p.Parser.GoString() + ")"
}

type SequenceParser []Parser

func NewSequenceParser(ps ...Parser) Parser {
	if len(ps) == 0 {
		panic("NewSequenceParser called with zero parsers")
	}
	if len(ps) == 1 {
		return ps[0]
	}
	return SequenceParser(ps)
}

func (p SequenceParser) Parse(bs []byte) StateSet {

	passups, rejects := NewStateSet(), NewStateSet()

	r := p[0].Parse(bs)

	for i, l := 0, r.Len(); i < l; i++ {

		s := r.Next()

		if s.Value == nil {

			cont := append([]Parser{s.Parser}, p[1:]...)

			passups.Add(
				NewState(
					s.Remnant,
					nil,
					NewSequenceParser(cont...),
				),
			)
			continue
		}

		if v, k := s.Value.(Accept); k {
			passups.Add(
				NewState(
					s.Remnant,
					nil,
					NewPrependParser(v.Value, NewSequenceParser(p[1:]...)),
				),
			)
			continue
		}

		rejects.Add(s)

	}

	if passups.Len() == 0 {
		return rejects
	}

	return passups
}

func (p SequenceParser) GoString() string {
	t := ""
	for _, q := range p {
		t += q.GoString() + "·"
	}
	t = t[:len(t)-2] // len("·") == 2
	return t
}

type MappingParser struct {
	Parser  Parser
	Mapping func(interface{}) interface{}
}

func NewMappingParser(p Parser, m func(interface{}) interface{}) Parser {
	return MappingParser{p, m}
}

func (p MappingParser) Parse(bs []byte) StateSet {

	r := p.Parser.Parse(bs)

	return r.Map(func(s State) State {

		if s.Value == nil {
			return NewState(
				s.Remnant,
				nil,
				NewMappingParser(s.Parser, p.Mapping),
			)
		}

		if v, k := s.Value.(Accept); k {
			return NewState(
				s.Remnant,
				Accept{(p.Mapping)(v.Value)},
				s.Parser,
			)
		}

		return s
	})
}

func (p MappingParser) GoString() string {
	return "M(" + p.Parser.GoString() + ")"
}

type RecursiveParser struct {
	Parser    Parser
	printFlag bool
}

func NewRecursiveParser(f func(Parser) Parser) Parser {
	p := &RecursiveParser{}
	p.Parser = f(p)
	return p
}

func (p *RecursiveParser) Parse(bs []byte) StateSet {
	return NewStateSet(NewState(bs, nil, p.Parser))
}

func (p *RecursiveParser) GoString() string {
	if p.printFlag {
		return "R"
	}
	p.printFlag = true
	t := p.Parser.GoString()
	p.printFlag = false
	return t
}

type KleeneParser struct {
	Parser Parser
}

func NewKleeneParser(p Parser) Parser {
	return KleeneParser{p}
}

func (p KleeneParser) Parse(bs []byte) StateSet {

	r := NewState(bs, nil, p.Parser).Reduce()

	return r.Map(func(s State) State {

		if s.Value == nil {
			return NewState(
				s.Remnant,
				nil,
				NewMaybeParser(
					NewSequenceParser(
						s.Parser,
						p,
					),
				),
			)
		}

		if v, k := s.Value.(Accept); k {
			return NewState(
				s.Remnant,
				nil,
				NewPrependParser(v.Value, p),
			)
		}

		return NewState(
			bs,
			Accept{bs[:0]},
			nil,
		)
	})
}

func (p KleeneParser) GoString() string {
	return p.Parser.GoString() + "*"
}

type MaybeParser struct {
	Parser Parser
}

func NewMaybeParser(p Parser) Parser {
	return MaybeParser{p}
}

func (p MaybeParser) Parse(bs []byte) StateSet {

	r := NewState(bs, nil, p.Parser).Reduce()

	return r.Map(func(s State) State {

		if s.Value == nil {
			return NewState(
				s.Remnant,
				nil,
				NewMaybeParser(s.Parser),
			)
		}

		if _, k := s.Value.(Accept); k {
			return s
		}

		return NewState(
			bs,
			Accept{bs[:0]},
			nil,
		)
	})
}

func (p MaybeParser) GoString() string {
	return p.Parser.GoString() + "?"
}
