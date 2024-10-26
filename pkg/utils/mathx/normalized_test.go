package mathx

import (
	"testing"
)

func TestNormalizationLog(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "Test",
			args: args{
				data: []float64{
					0.15, 0.3, 0.8, 0.5, 2.4, 3.6,
				},
			},
			want: []float64{
				0, 0.088438876, 0.3231815, 0.19166432, 0.78195053, 1,
			},
		},
		{
			name: "Maximal Test",
			args: args{
				data: []float64{
					0.15, 0.3, 0.8, 0.5, 2.4, 3.6, 20,
				},
			},
			want: []float64{
				0, 0.04220737, 0.15423808, 0.09147163, 0.37318516, 0.4772491, 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizationLog(tt.args.data)
			for i := range got {
				if !EqualWithThreshold(got[i], tt.want[i], 1e-5) {
					t.Errorf("NormalizationLog() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestNormalization(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "Test",
			args: args{
				data: []float64{
					0.15, 0.3, 0.8, 0.5, 2.4, 3.6,
				},
			},
			want: []float64{
				0, 0.043478265, 0.1884058, 0.10144928, 0.65217394, 1,
			},
		},
		{
			name: "Maximal Test",
			args: args{
				data: []float64{
					0.15, 0.3, 0.8, 0.5, 2.4, 3.6, 20,
				},
			},
			want: []float64{
				0, 0.007556675, 0.03274559, 0.01763224, 0.11335012, 0.17380351, 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Normalization(tt.args.data)
			for i := range got {
				if !EqualWithThreshold(got[i], tt.want[i], 1e-5) {
					t.Errorf("Normalization() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
