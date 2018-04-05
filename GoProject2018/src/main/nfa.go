package main

import (
	"fmt"
)

type state struct {
	symbol rune
	edge1  *state
	edge2  *state
} // End type

type nfa struct {
	initial *state
	accept  *state
} // End type

// Accepts user input
func GetInput() string {
	var input string
	fmt.Scan(&input)
	return input

}

// pofix regular expresion to nfa
func poregtonfa(pofix string) *nfa {
	nfastack := []*nfa{}

	for _, r := range pofix {
		switch r {
		case '.':

			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			frag1.accept.edge1 = frag2.initial

			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '|':
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		case '*':
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			intial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &intial, accept: &accept})

		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		} // End switch

	} // End for

	if len(nfastack) != 1 {
		fmt.Println("Uh Oh:", len(nfastack), nfastack)
	}

	return nfastack[0]

} // End func

func addState(l []*state, s *state, a *state) []*state {
	l = append(l, s)

	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		} // End if

	} // End if

	return l
} // End add state func

func pomatch(po string, s string) bool {
	ismatch := false
	ponfa := poregtonfa(po)

	current := []*state{}
	next := []*state{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	for _, r := range s {
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		} // End inner for
		current, next = next, []*state{}
	} // End

	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
		}

	}

	return ismatch

}

func main() {
	fmt.Print("Please enter a regular expression: ")

	var input1 = GetInput()

	fmt.Println("Postfix: ", intopost(input1))
	input1 = intopost(input1)

	fmt.Print("Enter a string: ")
	var input2 = GetInput()

	nfa := pomatch(input1, input2)
	fmt.Println("Match = ", nfa)

	fmt.Println(pomatch("ab.c*|", "c"))

}
