package main

import (
	"sort"
	"strings"
)

// Given 2 strings, S1 and S2,  return all substrings of S1 that are permutations of S2.
//
// E.g.
// S1: ABCDEC
// S2: DCE
// Result: [CDE,DEC]

func SubstringPerms(s1, s2 string) []string {
	// find all substrings of s1

	var result []string
	// for each substring, check if it is a permutation of s2
	for i := 0; i < len(s1); i++ {
		for j := i + 1; j <= len(s1); j++ {
			substr := s1[i:j]
			if isPermutation(substr, s2) {
				result = append(result, substr)
			}
		}
	}

	// if it is, add it to the result

	return result
}

func isPermutation(substr string, s2 string) bool {
	if len(substr) != len(s2) {
		return false
	}

	// sort substr and s2
	substr = sortString(substr)
	s2 = sortString(s2)

	// compare substr and s2
	return substr == s2
}

func sortString(substr string) string {
	s := strings.Split(substr, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
