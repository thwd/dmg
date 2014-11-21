package dmg

import (
	"fmt"
	"unicode/utf8"
)

type Accept struct {
	Value interface{}
}

type Reject struct {
	Value interface{}
}

type Parser interface {
	Parse([]byte) StateSet
	GoString() string
}

type State struct {
	Remnant []byte
	Value   interface{}
	Parser  Parser
}

func NewState(r []byte, v interface{}, p Parser) State {
	return State{r, v, p}
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

func (p RangeParser) Parse(bs []byte) StateSet {
	r, w := utf8.DecodeRune(bs)
	if w == 0 || r < p.Min || r > p.Max {
		return NewStateSet(NewState(bs, Reject{bs}, nil))
	}
	return NewStateSet(NewState(bs[w:], Accept{bs[:w]}, nil))
}

func (p RangeParser) Equals(o Parser) bool {
	q, k := o.(RangeParser)
	return (k && q.Min == p.Min && q.Max == p.Max)
}

func (p RangeParser) GoString() string {
	return "[" + string(p.Min) + "-" + string(p.Max) + "]"
}

type LiteralParser struct {
	Literal string
}

func NewLiteralParser(l string) Parser {
	return LiteralParser{l}
}

func (p LiteralParser) Parse(bs []byte) StateSet {
	if len(bs) < len(p.Literal) {
		return NewStateSet(NewState(bs, Reject{bs}, nil))
	}
	for i, l := 0, len(p.Literal); i < l; i++ {
		if bs[i] != p.Literal[i] {
			return NewStateSet(NewState(bs, Reject{bs}, nil))
		}
	}
	return NewStateSet(NewState(bs[len(p.Literal):], Accept{bs[:len(p.Literal)]}, nil))
}

func (p LiteralParser) GoString() string {
	return "\"" + p.Literal + "\""
}

type EpsilonParser struct{}

func NewEpsilonParser() Parser {
	return EpsilonParser{}
}

func (p EpsilonParser) Parse(bs []byte) StateSet {
	return NewStateSet(NewState(bs, Accept{bs[:0]}, nil))
}

func (p EpsilonParser) GoString() string {
	return "ε"
}
