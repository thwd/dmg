/*

Package dmg implements parser combinators to build simple and fast parsers.
The produced parsers are recursive-descent parsers with infinite lookahead,
aka. LL(k), but with support for left-recursive (in fact, any-recursive)
grammars. It also supports nullability (again, anywhere).

Additionally, dmg provides a set of convenience functions to provide common
building blocks for parsers such as the Kleene Closure and a Maybe-Parser,
equivalent in functionality to the POSIX regular expression operators `*`
and `?`, respectively.

A simple, left-recursive parser for arithmetic difference can be built as
follows:

    // number -> [0-9]+
    number := dmg.NewPlusParser(
        dmg.NewRangeParser('0', '9'),
    )

    // difference -> difference '-' number | number
    difference := dmg.NewRecursiveParser(func(self dmg.Parser) dmg.Parser {
        return dmg.NewAlternationParser(
            dmg.NewSequenceParser(
                self,
                dmg.NewLiteralParser("-"),
                number,
            ),
            number,
        )
    })

The above code, when evaluated to the end of a valid input will produce a
left-leaning concrete syntax tree. In order to build an abstract syntax tree,
dmg provides the MappingParser, which enables you to write "actions" to
manipulate matches. Following with the example:

    number := dmg.NewMappingParser(number, func(m interface{}) interface{} {
        return toInteger(m)
    })

This package is work-in-progress and will grow. Once it reaches a Beta-state
It'll get a grammar-tree optimizer which will hopefully eliminate any and all
back-tracking from the parser implementation, essentially guaranteeing an O(n)
time-complexity for unambiguous grammars, as well as scaling elegantly for
more complex grammars.

*/
package dmg
