package main

import "testing"

func Test_minChanges(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sample 1",
			args: args{
				s: "1001",
			},
			want: 2,
		},
		{
			name: "1 odd",
			args: args{
				s: "1011",
			},
			want: 1,
		},
		{
			name: "Three",
			args: args{
				s: "101001",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minChanges(tt.args.s); got != tt.want {
				t.Errorf("minChanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
