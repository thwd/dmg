package dmg

import (
	"fmt"
	"unicode/utf8"
)

type Remnant []byte

type Parser interface {
	Parse(Remnant) StateSet
}

type Value struct {
	Success bool
	Value   interface{}
}

func Accept(i interface{}) Value {
	return Value{true, i}
}

func Reject(i interface{}) Value {
	return Value{false, i}
}

type State struct {
	Final   bool
	Remnant Remnant
	Parser  Parser
	Value   Value
}

func NewFinalState(v Value, r Remnant) State {
	return State{true, r, nil, v}
}

func NewContinuedState(p Parser, r Remnant) State {
	return State{false, r, p, Value{false, nil}}
}

func (s State) Reduce() StateSet {
	return s.Parser.Parse(s.Remnant)
}

func (s State) GoString() string {
	t := "{\n"
	t += "\tRemnant: \"" + string(s.Remnant) + "\"\n"
	t += "\tParser:  " + fmt.Sprintf("%#v", s.Parser) + "\n"
	t += "\tValue:   " + fmt.Sprintf("%#v", s.Value) + "\n"
	t += "}\n"
	return t
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
			NewFinalState(Reject(bs), bs),
		)
	}
	return NewStateSet(
		NewFinalState(Accept(bs[:w]), bs[w:]),
	)
}

func (p RangeParser) Equals(o Parser) bool {
	q, k := o.(RangeParser)
	return (k && q.Min == p.Min && q.Max == p.Max)
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
			NewFinalState(Reject(bs), bs),
		)
	}
	for i, l := 0, len(p.Literal); i < l; i++ {
		if bs[i] != p.Literal[i] {
			return NewStateSet(
				NewFinalState(Reject(bs), bs),
			)
		}
	}
	w := len(p.Literal)
	return NewStateSet(
		NewFinalState(Accept(bs[:w]), bs[w:]),
	)
}

type EpsilonParser struct{}

func NewEpsilonParser() Parser {
	return EpsilonParser{}
}

func (p EpsilonParser) Parse(bs Remnant) StateSet {
	return NewStateSet(
		NewFinalState(Accept(bs[:0]), bs),
	)
}

type AnyParser struct{}

func NewAnyParser() Parser {
	return AnyParser{}
}

func (p AnyParser) Parse(bs Remnant) StateSet {
	if len(bs) == 0 {
		return NewStateSet(
			NewFinalState(Reject(bs), bs),
		)
	}
	return NewStateSet(
		NewFinalState(Accept(bs[:1]), bs[1:]),
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
			NewFinalState(Reject(bs), bs),
		)
	}
	return NewStateSet(
		NewFinalState(Accept(bs[:w]), bs[w:]),
	)
}
