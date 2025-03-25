package main

func validateParentheses(p string) bool {
	// there's equal amount of open and closed
	// for every open, there is a subsequent close
	// for every close, there is a preceding open
	c := 0
	for i := 0; i < len(p); i++ {
		if p[i] == '(' {
			c++
		}

		if p[i] == ')' {
			c--
		}

		if c < 0 {
			return false
		}
	}

	return c == 0
}
