package main

import "testing"
import "github.com/stretchr/testify/assert"

func Test_validateParentheses(t *testing.T) {
	tests := []struct {
		name string
		p    string
		want bool
	}{
		{
			name: "valid 1",
			p:    "()()",
			want: true,
		},
		{
			name: "valid 2",
			p:    "(())",
			want: true,
		},
		{
			name: "invalid 1",
			p:    "((())",
			want: false,
		},
		{
			name: "invalid 2",
			p:    "()())",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, validateParentheses(tt.p))
		})
	}
}
