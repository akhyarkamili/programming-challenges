package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_queue(t *testing.T) {
	q := queue{}
	assert.Equal(t, q.len(), 0)
	o := point{x: 1, y: 2}
	q.enqueue(o)
	assert.Equal(t, q.len(), 1)
	p := q.dequeue()
	assert.Equal(t, p, o)
	assert.Equal(t, q.len(), 0)
}

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "single island",
			args: args{
				[][]byte{
					[]byte("1"),
				},
			},
			want: 1,
		},
		{
			name: "no island",
			args: args{
				[][]byte{
					[]byte("0"),
				},
			},
			want: 0,
		},
		{
			name: "straight island",
			args: args{
				[][]byte{
					[]byte("110"),
				},
			},
			want: 1,
		},
		{
			name: "island",
			args: args{
				[][]byte{
					[]byte("000"),
					[]byte("001"),
				},
			},
			want: 1,
		},
		{
			name: "two islands, diagonal",
			args: args{
				[][]byte{
					[]byte("100"),
					[]byte("010"),
				},
			},
			want: 2,
		},
		{
			name: "two islands",
			args: args{
				[][]byte{
					[]byte("101"),
					[]byte("101"),
				},
			},
			want: 2,
		},
		{
			name: "two islands, single",
			args: args{
				[][]byte{
					[]byte("101"),
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numIslands(tt.args.grid), "numIslands(%v)", tt.args.grid)
		})
	}
}

func Test_isLand(t *testing.T) {
	type args struct {
		grid [][]byte
		p    point
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args{
				[][]byte{
					[]byte("100"),
				},
				point{0, 0},
			},
			true,
		},
		{
			args{
				[][]byte{
					[]byte("100"),
				},
				point{1, 0},
			},
			false,
		},
		{
			args{
				[][]byte{
					[]byte("100"),
					[]byte("001"),
				},
				point{2, 1},
			},
			true,
		},
		{
			args{
				[][]byte{
					[]byte("100"),
					[]byte("001"),
				},
				point{1, 1},
			},
			false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			assert.Equalf(t, tt.want, isLand(tt.args.grid, tt.args.p), "isLand(%v, %v)", tt.args.grid, tt.args.p)
		})
	}
}

func Test_isValidUnvisited(t *testing.T) {
	type args struct {
		visited map[point]bool
		n       point
		yLimit  int
		xLimit  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"single point not visited",
			args{map[point]bool{
				point{0, 0}: true,
			},
				point{0, 1},
				3,
				3,
			},
			true,
		},
		{
			"multiple points not visited",
			args{map[point]bool{
				point{0, 0}: true,
				point{2, 2}: true,
			},
				point{0, 1},
				3,
				3,
			},
			true,
		},
		{
			"multiple points visited",
			args{map[point]bool{
				point{0, 0}: true,
				point{2, 2}: true,
			},
				point{2, 2},
				3,
				3,
			},
			false,
		},
		{
			"single point out of bound",
			args{map[point]bool{
				point{0, 0}: true,
			},
				point{2, 0},
				1,
				1,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isValidUnvisited(tt.args.visited, tt.args.n, tt.args.yLimit, tt.args.xLimit), "isValidUnvisited(%v, %v, %v, %v)", tt.args.visited, tt.args.n, tt.args.yLimit, tt.args.xLimit)
		})
	}
}
