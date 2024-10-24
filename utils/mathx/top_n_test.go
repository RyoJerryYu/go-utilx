package mathx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceTopN(t *testing.T) {
	less := func(i, j int) bool { return i < j }
	greater := func(i, j int) bool { return i > j }
	type args struct {
		in   []int
		less func(int, int) bool
		N    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "短",
			args: args{
				[]int{1, 2, 3},
				less,
				5,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "greater",
			args: args{
				[]int{1, 2, 3, 8, 5, 9, 4},
				greater,
				5,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "less",
			args: args{
				[]int{1, 2, 3, 8, 5, 9, 4},
				less,
				5,
			},
			want: []int{8, 9, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SliceTopN(tt.args.in, tt.args.less, tt.args.N)
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}

func TestNewTopN(t *testing.T) {
	type data struct {
		value  int
		weight int
	}

	less := func(x, y data) bool {
		return x.weight < y.weight
	}

	greater := func(x, y data) bool {
		return x.weight > y.weight
	}

	type args struct {
		N    int
		less func(data, data) bool
		d    []data
	}
	tests := []struct {
		name string
		args args
		want []data
	}{
		{
			name: "最大topN",
			args: args{
				N:    2,
				less: less,
				d: []data{
					{6, 8}, {1, 2}, {5, 3},
				},
			},
			want: []data{
				{5, 3}, {6, 8},
			},
		},
		{
			name: "最小topN",
			args: args{
				N:    2,
				less: greater,
				d: []data{
					{5, 3}, {6, 8}, {1, 2},
				},
			},
			want: []data{
				{1, 2}, {5, 3},
			},
		},
		{
			name: "参数较少",
			args: args{
				N:    6,
				less: greater,
				d: []data{
					{1, 2}, {5, 3}, {6, 8},
				},
			},
			want: []data{
				{1, 2}, {5, 3}, {6, 8},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			topN := NewTopN(tt.args.N, tt.args.less)
			for _, item := range tt.args.d {
				topN.Push(item)
			}
			got := topN.Query()
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}
