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

    number := dmg.NewPlusParser(
        dmg.NewRangeParser('0', '9')
    )

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
I will work on a grammar-tree optimizer which will hopefully eliminate any
and all back-tracking from the parser implementation, essentially giving it a
best-case complexity of O(n), even for complicated grammars.

*/
package dmg
