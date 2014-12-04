package dmg

import (
	"unicode/utf8"
)

type Remnant []byte

type Parser interface {
	Parse(Remnant) StateSet
}

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

type LiteralParser struct {
	Literal string
}

func NewLiteralParser(l string) Parser {
	return LiteralParser{l}
}

func (p LiteralParser) Parse(bs Remnant) StateSet {
	if len(bs) < len(p.Literal) {
		return NewStateSet(
			Reject(bs, bs),
		)
	}
	for i, l := 0, len(p.Literal); i < l; i++ {
		if bs[i] != p.Literal[i] {
			return NewStateSet(
				Reject(bs, bs),
			)
		}
	}
	w := len(p.Literal)
	return NewStateSet(
		Accept(bs[:w], bs[w:]),
	)
}

type EpsilonParser struct{}

func NewEpsilonParser() Parser {
	return EpsilonParser{}
}

func (p EpsilonParser) Parse(bs Remnant) StateSet {
	return NewStateSet(
		Accept(bs[:0], bs),
	)
}

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

type NotRuneParser rune

func NewNotRuneParser(r rune) Parser {
	return NotRuneParser(r)
}

func (p NotRuneParser) Parse(bs Remnant) StateSet {
	r, w := utf8.DecodeRune(bs)
	if w == 0 || r == (rune)(p) {
		return NewStateSet(
			Reject(bs, bs),
		)
	}
	return NewStateSet(
		Accept(bs[:w], bs[w:]),
	)
}
