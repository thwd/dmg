package dmg

import (
	"unicode/utf8"
)

// A Remnant represents an input to a Parser
type Remnant []byte

type Parser interface {
	Parse(Remnant) StateSet
}

// RangeParser is a Parser that accepts an UTF-8 rune
// in the range [Min, Max] (inclusive)
type RangeParser struct {
	Min, Max rune
}

func NewRangeParser(min, max rune) Parser {
	return RangeParser{min, max}
}

func (p RangeParser) Parse(bs Remnant) StateSet {
	r, w := utf8.DecodeRune(bs)
	if w == 0 || r < p.Min || r > p.Max {
		return NewStateSet(
			Reject(bs, bs),
		)
	}
	return NewStateSet(
		Accept(bs[:w], bs[w:]),
	)
}

// LiteralParser is a Parser that accepts a given string
type LiteralParser string

func NewLiteralParser(l string) Parser {
	return LiteralParser(l)
}

func (p LiteralParser) Parse(bs Remnant) StateSet {
	if len(bs) < len(p) {
		return NewStateSet(
			Reject(bs, bs),
		)
	}
	for i, l := 0, len(p); i < l; i++ {
		if bs[i] != p[i] {
			return NewStateSet(
				Reject(bs, bs),
			)
		}
	}
	w := len(p)
	return NewStateSet(
		Accept(bs[:w], bs[w:]),
	)
}

// EpsilonParser is a Parser that always accepts the empty string
type EpsilonParser struct{}

func NewEpsilonParser() Parser {
	return EpsilonParser{}
}

func (p EpsilonParser) Parse(bs Remnant) StateSet {
	return NewStateSet(
		Accept(bs[:0], bs),
	)
}

// AnyParser is a Parser that accepts any one rune
type AnyParser struct{}

func NewAnyParser() Parser {
	return AnyParser{}
}

func (p AnyParser) Parse(bs Remnant) StateSet {
	if len(bs) == 0 {
		return NewStateSet(
			Reject(bs, bs),
		)
	}
	return NewStateSet(
		Accept(bs[:1], bs[1:]),
	)
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

func (p NotParser) Parse(r Remnant) StateSet {
	return p.Parser.Parse(r).Map(func(s State) State {
		if s.Continued() {
			return s
		}
		if s.Accepted() {
			return Reject(s.Value, s.Remnant)
		}
		return Continue(NewAnyParser(), r)
	})
}
