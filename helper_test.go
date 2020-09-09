package dbhelper

import "testing"

func TestPageCount(t *testing.T) {
	type args struct {
		total    int
		pagesize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{15, 10}, 2},
		{"test2", args{20, 10}, 2},
		{"test3", args{21, 10}, 3},
		{"test4", args{5, 10}, 1},
		{"test5", args{0, 10}, 0},
		{"test6", args{5, 0}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PageCount(tt.args.total, tt.args.pagesize); got != tt.want {
				t.Errorf("PageCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
