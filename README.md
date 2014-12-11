## dmg
##### Simple and Fast (Generalized) Parsing for Go

Package `dmg` (pronounced ["demagogue"](http://en.wikipedia.org/wiki/Demagogue)) is a parser combinator library for the Go programming language. It is fast, has full support for left-recursive and right-nullable grammars and stands out for its simplicity in design.

Full documentation can be found on http://godoc.org/github.com/thwd/dmg.

It's not quite ready for prime-time yet but it works, and works well so far. There will be more added features in the near future. I'm working on finishing a stable API and will promote it to beta then.

This project contains some original research of mine, built upon the following articles and theses:
 * [Matt Might's "Yacc is dead"](http://matt.might.net/articles/parsing-with-derivatives/)
 * [Russ Cox's "Yacc is not dead"](http://research.swtch.com/yaccalive)
 * [Janusz Brzozowksi's "Derivatives of regular expressions"](http://dl.acm.org/citation.cfm?doid=321239.321249)
 * [Paul Lickman's "Parsing with fixedpoints"](https://sites.google.com/a/lickman.com/paul-lickman/)

An article on said research coming soon.
