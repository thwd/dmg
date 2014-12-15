package dmg

import (
	"unicode/utf8"
)

// A Remnant represents an input to a Parser
type Remnant []byte

type Parser interface {
	Parse(Remnant, chan State)
}

// RangeParser is a Parser that accepts an UTF-8 rune
// in the range [Min, Max] (inclusive)
type RangeParser struct {
	Min, Max rune
}

func NewRangeParser(min, max rune) Parser {
	return RangeParser{min, max}
}

func (p RangeParser) Parse(r Remnant, c chan State) {

	a, w := utf8.DecodeRune(r)

	if w == 0 || a < p.Min || a > p.Max {
		c <- Reject(nil, r)
		return
	}

	c <- Accept(r[:w], r[w:])
}

// LiteralParser is a Parser that accepts a given string
type LiteralParser string

func NewLiteralParser(l string) Parser {
	return LiteralParser(l)
}

func (p LiteralParser) Parse(r Remnant, c chan State) {

	if len(r) < len(p) {
		c <- Reject(nil, r)
		return
	}

	for i, l := 0, len(p); i < l; i++ {
		if r[i] != p[i] {
			c <- Reject(nil, r)
			return
		}
	}

	c <- Accept(r[:len(p)], r[len(p):])

}

// EpsilonParser is a Parser that always accepts the empty string
type EpsilonParser struct{}

func NewEpsilonParser() Parser {
	return EpsilonParser{}
}

func (p EpsilonParser) Parse(r Remnant, c chan State) {
	c <- Accept(r[:0], r)
}

// AnyParser is a Parser that accepts any one rune
type AnyParser struct{}

func NewAnyParser() Parser {
	return AnyParser{}
}

func (p AnyParser) Parse(r Remnant, c chan State) {

	if len(r) == 0 {
		c <- Reject(nil, r)
		return
	}

	c <- Accept(r[:1], r[1:])

}

// NotParser is a Parser that accepts any one rune from an input that
// is rejected by a given Parser. It rejects anything that said Parser
// accepts, without consuming any input.
type NotParser struct {
	Parser Parser
}

func NewNotParser(p Parser) Parser {
	return NotParser{p}
}

func (p NotParser) Parse(r Remnant, c chan State) {

	d, f := make(chan State), make(chan struct{})

	go func() {

		for s := range d {

			if s.Continued() {
				c <- s
				continue
			}

			if s.Accepted() {
				c <- Reject(p.Parser, s.Remnant)
				continue
			}

			NewAnyParser().Parse(r, c)
		}

		f <- struct{}{}

	}()

	p.Parser.Parse(r, d)

	close(d)

	<-f
}
