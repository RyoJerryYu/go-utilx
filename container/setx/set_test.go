package setx

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := []int{3, 4, 5, 6, 7, 8}

	setA := FromSlice(sliceA)
	setB := FromSlice(sliceB)

	assert.Len(t, setA, 5)
	assert.Len(t, setB, 6)

	union := Union(setA, setB)
	assert.Len(t, union, 8)
	for i := 1; i <= 8; i++ {
		assert.Contains(t, union, i)
	}

	intersect := Intersect(setA, setB)
	assert.Len(t, intersect, 3)
	for i := 3; i <= 5; i++ {
		assert.Contains(t, intersect, i)
	}

	subtract := Subtract(setA, setB)
	assert.Len(t, subtract, 2)
	for i := 1; i <= 2; i++ {
		assert.Contains(t, subtract, i)
	}
}

func TestIntersectSlice(t *testing.T) {
	type args struct {
		a map[int64]struct{}
		b []int64
	}
	tests := []struct {
		name string
		args args
		want map[int64]struct{}
	}{
		{
			name: "普通测试",
			args: args{
				a: map[int64]struct{}{
					2: {},
					5: {},
				},
				b: []int64{0, 2, 3, 4},
			},
			want: map[int64]struct{}{
				2: {},
			},
		},
		{
			name: "空-1",
			args: args{
				a: map[int64]struct{}{},
				b: []int64{0, 2, 3, 4},
			},
			want: map[int64]struct{}{},
		},
		{
			name: "空-2",
			args: args{
				a: map[int64]struct{}{
					2: {},
					5: {},
				},
				b: []int64{},
			},
			want: map[int64]struct{}{},
		},
		{
			name: "空-3",
			args: args{
				a: map[int64]struct{}{
					1: {},
					5: {},
				},
				b: []int64{0, 2, 3, 4},
			},
			want: map[int64]struct{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntersectSlice(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectSetWithSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
