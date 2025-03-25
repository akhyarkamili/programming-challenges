package main

import (
	"reflect"
	"testing"
)

func TestSubstringPerms(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test case 1",
			args: args{
				s1: "ABCDEC",
				s2: "DCE",
			},
			want: []string{"CDE", "DEC"},
		},
		{
			name: "s1 is empty",
			args: args{
				s1: "",
				s2: "DCE",
			},
			want: []string{},
		},
		{
			name: "s2 is empty",
			args: args{
				s1: "ABD",
				s2: "",
			},
			want: []string{},
		},
		{
			name: "s2 has duplicate characters",
			args: args{
				s1: "ABDCDD",
				s2: "DDC",
			},
			want: []string{"DCD", "CDD"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubstringPerms(tt.args.s1, tt.args.s2); !(reflect.DeepEqual(got, tt.want) || (len(got) == 0 && len(tt.want) == 0)) {
				t.Errorf("SubstringPerms() = %v, want %v", got, tt.want)
			}
		})
	}
}
