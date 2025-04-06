package main

type token struct {
	char      uint8
	repeating bool
	endIdx    int
}

func nextToken(s string, idx int) token {
	maxLen := len(s)
	if idx >= maxLen {
		return token{}
	}

	var char = s[idx]
	var t = token{
		char:   char,
		endIdx: idx,
	}

	if idx < len(s)-1 && s[idx+1] == 42 {
		t.repeating = true
		t.endIdx++
	}
	return t
}

// isTokenMatch returns whether character matches the parsing token
func isTokenMatch(c string, t *token) bool {
	if t.char == 46 {
		return true
	}

	if c == "" {
		return t.repeating
	}

	return t.char == c[0]
}

type input struct {
	s string
	p string
}

func isMatchMemoized(s string, p string, memo map[input]bool) bool {
	key := input{s, p}
	var result bool
	if result, ok := memo[key]; ok {
		return result
	}

	if len(s) == 0 && len(p) == 0 {
		memo[key] = true
		return true
	}
	if len(p) == 0 && len(s) != 0 {
		memo[key] = false
		return false
	}

	t := nextToken(p, 0)
	if t.repeating {
		if isTokenMatch(s, &t) {
			result = (len(s) > 0 && isMatchMemoized(s[1:], p, memo)) || isMatchMemoized(s, p[t.endIdx+1:], memo)
			memo[key] = result
			return result
		}

		result = isMatchMemoized(s, p[t.endIdx+1:], memo)
		memo[key] = result
		return result
	}

	result = len(s) > 0 && isTokenMatch(s, &t) && isMatchMemoized(s[1:], p[t.endIdx+1:], memo)
	memo[key] = result
	return result
}

func isMatch(s string, p string) bool {
	return isMatchMemoized(s, p, make(map[input]bool))
}
