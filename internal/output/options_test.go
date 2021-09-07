package output

import (
	"reflect"
	"testing"

	"github.com/ffhuo/swaggotest/internal/models"
	"github.com/ffhuo/swaggotest/internal/render"
)

func TestOptions_Process(t *testing.T) {
	type fields struct {
		PrintInputs    bool
		Subtests       bool
		Parallel       bool
		Named          bool
		Template       string
		TemplateDir    string
		TemplateParams map[string]interface{}
		TemplateData   [][]byte
		render         *render.Render
	}
	type args struct {
		head  *models.Header
		funcs []*models.Function
	}

	// receiver := new(models.Receiver)
	// receiver.Field = new(models.Field)
	// receiver.Name = ""
	// receiver.Type = &models.Expression{
	// 	Value:      "AccountDiscountForm",
	// 	IsStar:     true,
	// 	IsVariadic: false,
	// 	IsWriter:   false,
	// 	Underlying: "",
	// }
	// receiver.Fields = []*models.Field{
	// 	{
	// 		Name: "aa",
	// 		Type: &models.Expression{
	// 			Value:      "int",
	// 			IsStar:     false,
	// 			IsVariadic: false,
	// 			IsWriter:   false,
	// 			Underlying: "",
	// 		},
	// 		Index: 0,
	// 	},
	// 	{
	// 		Name: "bb",
	// 		Type: &models.Expression{
	// 			Value:      "int",
	// 			IsStar:     false,
	// 			IsVariadic: false,
	// 			IsWriter:   false,
	// 			Underlying: "",
	// 		},
	// 		Index: 0,
	// 	},
	// }

	funcSetting := &models.Function{
		Name:       "ls",
		IsExported: false,
		// Receiver:   receiver,
		Parameters: []*models.Field{
			{
				Name: "a",
				Type: &models.Expression{
					Value:      "int",
					IsStar:     false,
					IsVariadic: false,
					IsWriter:   false,
					Underlying: "",
				},
				Index: 0,
			},
			{
				Name: "b",
				Type: &models.Expression{
					Value:      "int",
					IsStar:     false,
					IsVariadic: false,
					IsWriter:   false,
					Underlying: "fields",
				},
				Index: 1,
			},
			{
				Name: "form",
				Type: &models.Expression{
					Value:      "int",
					IsStar:     false,
					IsVariadic: false,
					IsWriter:   false,
					Underlying: "",
				},
				Index: 2,
			},
		},
		Results: []*models.Field{
			{
				Name: "",
				Type: &models.Expression{
					Value:      "int",
					IsStar:     true,
					IsVariadic: true,
					IsWriter:   false,
					Underlying: "",
				},
				Index: 0,
			},
		},
		ReturnsError: false,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test_Process",
			fields: fields{
				PrintInputs:    true,
				Subtests:       false,
				Parallel:       true,
				Named:          true,
				Template:       "",
				TemplateDir:    "",
				TemplateParams: make(map[string]interface{}),
			},
			args: args{
				head: &models.Header{
					Comments: []string{},
					Package:  "swaggo",
				},
				funcs: []*models.Function{
					funcSetting,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				PrintInputs:    tt.fields.PrintInputs,
				Subtests:       tt.fields.Subtests,
				Parallel:       tt.fields.Parallel,
				Named:          tt.fields.Named,
				Template:       tt.fields.Template,
				TemplateDir:    tt.fields.TemplateDir,
				TemplateParams: tt.fields.TemplateParams,
				TemplateData:   tt.fields.TemplateData,
				render:         tt.fields.render,
			}
			got, err := o.Process(tt.args.head, tt.args.funcs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Options.Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Options.Process() = %s, want %v", string(got), tt.want)
			}
		})
	}
}
