package dmg

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
