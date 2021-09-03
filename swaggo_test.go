package swaggo

import (
	"testing"

	"github.com/ffhuo/swaggotest/internal/models"
)

func TestReadDoc(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.SwaggerData
		wantErr bool
	}{
		{
			name: "test_read_swag_doc",
			args: args{
				fileName: "../swagger.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// got, err := ReadDoc(tt.args.fileName)
			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("ReadDoc() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			// b, _ := json.Marshal(got)
			// t.Errorf("ReadDoc() = %s, want %v", string(b), tt.want)
		})
	}
}
