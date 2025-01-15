package httpx

import (
	"testing"
)

func TestBuildURL(t *testing.T) {
	type args struct {
		baseURL     string
		queryParams map[string][]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				baseURL: "www.abc.com",
				queryParams: map[string][]string{
					"b": {"2"},
				},
			},
			want:    "www.abc.com?b=2",
			wantErr: false,
		},
		{
			name: "simple-1",
			args: args{
				baseURL: "www.abc.com?a=1",
				queryParams: map[string][]string{
					"b": {"2"},
				},
			},
			want:    "www.abc.com?a=1&b=2",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildURL(tt.args.baseURL, tt.args.queryParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BuildURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
