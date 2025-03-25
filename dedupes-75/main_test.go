package main

import (
	"os"
	"slices"
	"testing"
)

func Test_listFiles(t *testing.T) {
	type args struct {
		fileSetup func(dir string)
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Simple flat dir",
			args: args{
				func(dir string) {
					f1, _ := os.Create(dir + "/file1.txt")
					f2, _ := os.Create(dir + "/file2.txt")
					f3, _ := os.Create(dir + "/file3.txt")

					//
					_, _ = f1.Write([]byte("Hello, world"))
					_, _ = f2.Write([]byte("Hello, world!"))
					_, _ = f3.Write([]byte("Hello, world"))
				},
			},
			want: []string{"file1.txt", "file2.txt", "file3.txt"},
		},
		{
			name: "Simple nested dir",
			args: args{
				func(dir string) {
					_ = os.Mkdir(dir+"/dir1", 0755)
					_ = os.Mkdir(dir+"/dir2", 0755)
					_ = os.Mkdir(dir+"/dir3", 0755)

					f1, _ := os.Create(dir + "/file1.txt")
					f2, _ := os.Create(dir + "/dir2/file2.txt")
					f3, _ := os.Create(dir + "/dir3/file3.txt")

					//
					_, _ = f1.Write([]byte("Hello, world"))
					_, _ = f2.Write([]byte("Hello, world!"))
					_, _ = f3.Write([]byte("Hello, world"))
				},
			},
			want: []string{"file1.txt", "dir2/file2.txt", "dir3/file3.txt"},
		},
		{
			name: "Multi-level nested dir",
			args: args{
				func(dir string) {
					_ = os.Mkdir(dir+"/dir1", 0755)
					_ = os.Mkdir(dir+"/dir1/dir2", 0755)
					_ = os.Mkdir(dir+"/dir3", 0755)

					f1, _ := os.Create(dir + "/file1.txt")
					f2, _ := os.Create(dir + "/dir1/dir2/file2.txt")
					f3, _ := os.Create(dir + "/dir3/file3.txt")

					//
					_, _ = f1.Write([]byte("Hello, world"))
					_, _ = f2.Write([]byte("Hello, world!"))
					_, _ = f3.Write([]byte("Hello, world"))
				},
			},
			want: []string{"file1.txt", "dir1/dir2/file2.txt", "dir3/file3.txt"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			tt.args.fileSetup(dir)
			res := listFiles(dir)
			setFilePath := make(map[string]struct{})

			for _, file := range res {
				setFilePath[file.relativePath] = struct{}{}
			}

			if len(setFilePath) != len(tt.want) {
				t.Errorf("listFiles() = %v, want %v", setFilePath, tt.want)
			}

			for _, file := range tt.want {
				if _, ok := setFilePath[file]; !ok {
					t.Errorf("listFiles() = %v, want %v", setFilePath, tt.want)
				}
			}
		})
	}
}

func Test_identifyDuplicates(t *testing.T) {
	type args struct {
		fileSetup func(dir string)
	}
	tests := []struct {
		name string
		args args
		want map[int64][]string
	}{
		{
			name: "Two same size",
			args: args{
				func(dir string) {
					f1, _ := os.Create(dir + "/file1.txt")
					f2, _ := os.Create(dir + "/file2.txt")
					f3, _ := os.Create(dir + "/file3.txt")

					_, _ = f1.Write([]byte("Hello, world"))
					_, _ = f2.Write([]byte("Hello, world!"))
					_, _ = f3.Write([]byte("Hello, world"))
				},
			},
			want: map[int64][]string{ // compare only the relativePath
				12: {
					"file1.txt",
					"file3.txt",
				},
				13: {
					"file2.txt",
				},
			},
		},
		{
			name: "No same size",
			args: args{
				func(dir string) {
					_ = os.Mkdir(dir+"/dir1", 0755)
					_ = os.Mkdir(dir+"/dir2", 0755)
					_ = os.Mkdir(dir+"/dir3", 0755)

					f1, _ := os.Create(dir + "/file1.txt")
					f2, _ := os.Create(dir + "/dir2/file2.txt")
					f3, _ := os.Create(dir + "/dir3/file3.txt")

					_, _ = f1.Write([]byte("Hello, worl"))
					_, _ = f2.Write([]byte("Hello, world!"))
					_, _ = f3.Write([]byte("Hello, world"))
				},
			},
			want: map[int64][]string{
				11: {
					"file1.txt",
				},
				12: {
					"dir3/file3.txt",
				},
				13: {
					"dir2/file2.txt",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			tt.args.fileSetup(dir)
			files := listFiles(dir)
			got := identifyDuplicates(files, dir)
			for k, v := range got {
				if !setEqual(tt.want[k], v, func(file fileInfo) string {
					return file.relativePath
				}) {
					t.Errorf("identifyDuplicates() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_setEqual(t *testing.T) {
	type args struct {
		a         []int
		b         []string
		transform func(string) int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Equal",
			args: args{
				a: []int{1, 2, 3},
				b: []string{"1", "2", "3"},
				transform: func(s string) int {
					return int(s[0] - '0')
				},
			},
			want: true,
		},
		{
			name: "Not equal",
			args: args{
				a: []int{1, 2, 3},
				b: []string{"1", "2", "4"},
				transform: func(s string) int {
					return int(s[0] - '0')
				},
			},
			want: false,
		},

		{
			name: "Not equal",
			args: args{
				a: []int{1, 2, 4, 5},
				b: []string{"1", "2", "4"},
				transform: func(s string) int {
					return int(s[0] - '0')
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setEqual(tt.args.a, tt.args.b, tt.args.transform); got != tt.want {
				t.Errorf("setEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func setEqual[T comparable, V any](a []T, b []V, transform func(V) T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range b {
		if !slices.Contains(a, transform(b[i])) {
			return false
		}
	}

	return true
}
