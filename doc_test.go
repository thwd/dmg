package dmg

import (
	"fmt"
	"github.com/thwd/dmg"
)

func Example() {
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

	treec := dmg.Dissect(difference, dmg.Remnant("10-7-1"))

	// print all successful matches
	for {

		state := <-treec

		if state.Rejected() {
			continue // don't care
		}

		fmt.Printf("%#v\n", dmg.MatchToString(state.Value))

		if len(state.Remnant) == 0 {
			break // we're done
		}
	}
}
