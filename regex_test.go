package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"matching exact characters",
			args{
				s: "abcde",
				p: "abcde",
			},
			true,
		},
		{
			"not matching partial characters",
			args{
				s: "abcde",
				p: "abc",
			},
			false,
		},
		{
			"any character match",
			args{
				s: "aaa",
				p: "aa.",
			},
			true,
		},
		{
			"repeating character match",
			args{
				s: "aaaaa",
				p: "a*",
			},
			true,
		},
		{
			"repeating character match 2",
			args{
				s: "baaaa",
				p: "ba*",
			},
			true,
		},
		{
			"repeating character match middle",
			args{
				s: "baaad",
				p: "ba*d",
			},
			true,
		},
		{
			"repeating character match empty",
			args{
				s: "ba",
				p: "c*ba",
			},
			true,
		},
		{
			"repeating character match empty 2",
			args{
				s: "bd",
				p: "c*ba*d",
			},
			true,
		},
		{
			"repeating character followed by something else",
			args{
				s: "cccbd",
				p: ".*c",
			},
			false,
		},
		{
			"repeating character followed by something else 2",
			args{
				s: "ccc",
				p: "c*c",
			},
			true,
		},
		{
			"longer pattern",
			args{
				s: "bd",
				p: "c*ba*deee",
			},
			false,
		},
		{
			"longer string",
			args{
				s: "bdeeee",
				p: "c*ba*deee",
			},
			false,
		},
		{
			"mixed with wildcard and repeating",
			args{
				s: "baaadd",
				p: "ba*d.",
			},
			true,
		},
		{
			"mixed with wildcard and repeating with non-repeating correctness",
			args{
				s: "aab",
				p: ".*a*",
			},
			true,
		},
		{
			"wildcard and repeating",
			args{
				s: "baaaddasdfdsfc",
				p: ".*c",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isMatch(tt.args.s, tt.args.p), "isMatch(%v, %v)", tt.args.s, tt.args.p)
		})
	}
}

func Test_isTokenMatch(t *testing.T) {
	type args struct {
		s string
		t token
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"exact match",
			args{
				"c",
				token{
					99,
					false,
					1,
				},
			},
			true,
		},
		{
			"exact match repeating",
			args{
				"c",
				token{
					99,
					true,
					1,
				},
			},
			true,
		},
		{
			"no match repeating",
			args{
				"c",
				token{
					98,
					true,
					1,
				},
			},
			true,
		},
		{
			"exact match false",
			args{
				"c",
				token{
					98,
					false,
					1,
				},
			},
			false,
		},
		{
			"any match",
			args{
				"c",
				token{
					46,
					false,
					1,
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isTokenMatch(tt.args.s, &tt.args.t), "isTokenMatch(%v, %v)", tt.args.s, tt.args.t)
		})
	}
}

func Test_nextToken(t *testing.T) {
	type args struct {
		s   string
		idx int
	}
	tests := []struct {
		name string
		args args
		want token
	}{
		{
			"empty string",
			args{
				"",
				0,
			},
			token{},
		},
		{
			"single char a",
			args{
				"a",
				0,
			},
			token{
				endIdx:    0,
				char:      97,
				repeating: false,
			},
		},
		{
			"single char .",
			args{
				".",
				0,
			},
			token{
				endIdx:    0,
				char:      46,
				repeating: false,
			},
		},
		{
			"single char in many chars",
			args{
				"aabaaa",
				2,
			},
			token{
				endIdx:    2,
				char:      98,
				repeating: false,
			},
		},
		{
			"b* in many chars",
			args{
				"aab*aa",
				2,
			},
			token{
				endIdx:    3,
				char:      98,
				repeating: true,
			},
		},
		{
			"b* as last char",
			args{
				"aab*",
				2,
			},
			token{
				endIdx:    3,
				char:      98,
				repeating: true,
			},
		},
		{
			"single repeating",
			args{
				"a*",
				0,
			},
			token{
				endIdx:    1,
				char:      97,
				repeating: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, nextToken(tt.args.s, tt.args.idx), "nextToken(%v, %v)", tt.args.s, tt.args.idx)
		})
	}
}

func useMap(i map[int]bool) {
	i[1] = true
}

func TestMap(t *testing.T) {
	m := make(map[int]bool)
	useMap(m)
	assert.Equal(t, true, m[1])
	assert.Equal(t, false, m[2])
}
