package dmg

// PrependParser is a Parser used by SequenceParser to prepend a value
// to the Value of each accepted State returned by a given Parser
type PrependParser struct {
	Prepend interface{}
	Parser  Parser
}

func NewPrependParser(r interface{}, p Parser) Parser {
	return PrependParser{r, p}
}

func (p PrependParser) Parse(r Remnant, c chan State) {

	d, f := make(chan State), make(chan struct{})

	go func() {
		for s := range d {

			if s.Continued() {
				c <- Continue(
					NewPrependParser(p.Prepend, s.Parser),
					s.Remnant,
				)
				continue
			}

			if s.Accepted() {
				c <- Accept(
					[2]interface{}{p.Prepend, s.Value},
					s.Remnant,
				)
				continue
			}

			c <- Reject(
				[2]interface{}{p.Prepend, s.Value},
				s.Remnant,
			)
		}

		f <- struct{}{}
	}()

	p.Parser.Parse(r, d)

	close(d)

	<-f
}
