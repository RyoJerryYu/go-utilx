package slicex

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeduplicatedSlice(t *testing.T) {
	sourceA := []int{1, 2, 3, 4, 5, 5, 5, 5}
	sourceB := []int{3, 3, 4, 5, 6, 7, 8}

	sliceA := Deduplicate(sourceA)
	sliceB := Deduplicate(sourceB)

	assert.Len(t, sliceA, 5)
	assert.Len(t, sliceB, 6)

	union := Union(sliceA, sliceB)
	assert.Len(t, union, 8)
	for i := 1; i <= 8; i++ {
		assert.Equal(t, i, union[i-1])
	}

	intersect := Intersect(sliceA, sliceB)
	assert.Len(t, intersect, 3)
	for i := 3; i <= 5; i++ {
		assert.Equal(t, i, intersect[i-3])
	}

	subtract := Subtract(sliceA, sliceB)
	assert.Len(t, subtract, 2)
	for i := 1; i <= 2; i++ {
		assert.Equal(t, i, subtract[i-1])
	}
}

func TestMergeAll(t *testing.T) {
	type args struct {
		slices [][]int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "Merge all",
			args: args{
				slices: [][]int64{
					{4, 5, 6},
					{5, 6, 7, 8},
					{8, 9},
				},
			},
			want: []int64{4, 5, 6, 5, 6, 7, 8, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeAll(tt.args.slices...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeduplicateBack(t *testing.T) {
	sourceA := []int{1, 2, 3, 4, 2, 5}
	sourceB := []int{3, 4, 5, 7, 6, 7, 8}

	sliceA := DeduplicateBack(sourceA)
	sliceB := DeduplicateBack(sourceB)

	assert.Equal(t, []int{1, 3, 4, 2, 5}, sliceA)
	assert.Equal(t, []int{3, 4, 5, 6, 7, 8}, sliceB)
}
