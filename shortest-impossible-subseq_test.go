package main

import "testing"

func Test_shortestSequence(t *testing.T) {
	type args struct {
		rolls []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no 1, 2, but there is 3",
			args: args{
				rolls: []int{4, 2, 1, 2, 3, 3, 2, 4, 1},
				k:     4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestSequence(tt.args.rolls, tt.args.k); got != tt.want {
				t.Errorf("shortestSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
