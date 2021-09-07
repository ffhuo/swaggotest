package swaggo

import (
	"reflect"
	"testing"
)

func TestGenerateTests(t *testing.T) {
	type args struct {
		opt *Options
	}
	tests := []struct {
		name    string
		args    args
		want    []*GeneratedTest
		wantErr bool
	}{
		{
			name: "test_option",
			args: args{
				opt: &Options{
					Path: "./swagger.json",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateTests(tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateTests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%s", string(got[0].Output))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateTests() = %v, want %v", got, tt.want)
			}
		})
	}
}
