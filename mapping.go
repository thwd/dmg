package dmg

// MappingParser is a Parser that matches a given Parser against a Remnant
// and mutates the Values of all accepted States through a mapping function.
type MappingParser struct {
	Parser  Parser
	Mapping func(interface{}) interface{}
}

func NewMappingParser(p Parser, m func(interface{}) interface{}) Parser {
	return MappingParser{p, m}
}

// Parse delegates to the MappingParser's underlying parser and maps all
// matched states' values through the MappingParser's mapping function.
func (p MappingParser) Parse(r Remnant, c chan State) {

	d, f := make(chan State), make(chan struct{})

	go func() {
		for s := range d {

			if s.Continued() {
				c <- Continue(
					NewMappingParser(s.Parser, p.Mapping),
					s.Remnant,
				)
				continue
			}

			if s.Accepted() {
				c <- Accept(
					(p.Mapping)(s.Value),
					s.Remnant,
				)
				continue
			}

			c <- s
		}

		f <- struct{}{}
	}()

	p.Parser.Parse(r, d)

	close(d)

	<-f
}
