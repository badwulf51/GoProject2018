package main

import (
	"fmt"
)

func intopost(infix string) string {
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}

	pofix, s := []rune{}, []rune{}

	for _, r := range infix {
		switch {
		case r == '(':
			s = append(s, r)

		case r == ')':
			for s[len(s)-1] != '(' {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			} // End For
			s = s[:len(s)-1]
		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {

				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]

			} // End For
			s = append(s, r)
		default:
			pofix = append(pofix, r)

		} // End Switch

	} // End For

	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}

	return string(pofix)
}

func mains() {
	// Answer: ab.c*
	fmt.Println("infix:   ", "a.b.c*")
	fmt.Println("postfix  ", intopost("a.b.c*"))
	// Answer: abd|.*
	fmt.Println("infix:   ", "(a.(b|d))*")
	fmt.Println("postfix  ", intopost("(a.(b|d))*"))
	// Answer: abd|.c*
	fmt.Println("infix:   ", "a.(b|d).c*")
	fmt.Println("postfix  ", intopost("a.(b|d).c*"))
	// Answer abb.+.c.
	fmt.Println("infix:   ", "a.(b.b)+.c")
	fmt.Println("postfix  ", intopost("a.(b.b)+.c"))

}
