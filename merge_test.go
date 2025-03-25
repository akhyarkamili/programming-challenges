package main

import "testing"

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sample",
			args: args{
				nums1: []int{1, 2, 2, 0, 0, 0},
				nums2: []int{1, 3, 6},
				m:     3,
				n:     3,
			},
			want: []int{1, 1, 2, 2, 3, 6},
		},
		{
			name: "nums2 all smaller",
			args: args{
				nums1: []int{7, 8, 9, 0, 0, 0},
				nums2: []int{1, 3, 6},
				m:     3,
				n:     3,
			},
			want: []int{1, 3, 6, 7, 8, 9},
		},
		{
			name: "nums1 all smaller",
			args: args{
				nums1: []int{0, 0, 0, 0, 0, 0},
				nums2: []int{1, 3, 6},
				m:     3,
				n:     3,
			},
			want: []int{0, 0, 0, 1, 3, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			for i := range tt.args.nums1 {
				if tt.want[i] != tt.args.nums1[i] {
					t.Errorf("Failed case %s, at index: %d, nums1: %v, want: %v", tt.name, i, tt.args.nums1, tt.want)
				}
			}
		})
	}
}
