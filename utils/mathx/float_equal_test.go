package mathx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloatEqual(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{1, 1}, true},
		{"2", args{1, 0.9}, false},
		{"3", args{1, 1.0000000000}, true},
		{"4", args{1, 1.0000000001}, true},
		{"5", args{0, 1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Equal(tt.args.a, tt.args.b))
		})
	}
}
