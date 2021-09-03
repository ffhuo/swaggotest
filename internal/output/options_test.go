package output

import (
	"reflect"
	"testing"

	"github.com/ffhuo/swaggotest/internal/models"
	"github.com/ffhuo/swaggotest/internal/render"
)

func TestOptions_providesTemplateData(t *testing.T) {
	tests := []struct {
		name    string
		otpions *Options
		want    bool
	}{
		{"Opt is nil", nil, false},
		{"Opt is empty", &Options{}, false},
		{"TemplateData is nil", &Options{TemplateData: nil}, false},
		{"TemplateData is empty", &Options{TemplateData: [][]byte{}}, false},
		{"TemplateData is OK", &Options{TemplateData: [][]byte{[]byte("ok")}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.otpions.providesTemplateData(); got != tt.want {
				t.Errorf("Options.isProvidesTemplateData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_providesTemplate(t *testing.T) {
	tests := []struct {
		name    string
		otpions *Options
		want    bool
	}{
		{"Opt is nil", nil, false},
		{"Opt is empty", &Options{}, false},
		{"Template is empty (implicit_zero_val)", &Options{Template: ""}, false},
		{"Template is OK", &Options{Template: "testify"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.otpions.providesTemplate(); got != tt.want {
				t.Errorf("Options.isProvidesTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_providesTemplateDir(t *testing.T) {
	tests := []struct {
		name    string
		otpions *Options
		want    bool
	}{
		{"Opt is nil", nil, false},
		{"Opt is empty", &Options{}, false},
		{"Template is empty", &Options{TemplateDir: ""}, false},
		{"Template is OK", &Options{TemplateDir: "testify"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.otpions.providesTemplateDir(); got != tt.want {
				t.Errorf("Options.isProvidesTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

	funcSetting := &models.Function{
		Name:       "ls",
		IsExported: false,
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
					Underlying: "",
				},
				Index: 0,
			},
		},
		Results: []*models.Field{
			{
				Name: "",
				Type: &models.Expression{
					Value:      "int",
					IsStar:     false,
					IsVariadic: false,
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