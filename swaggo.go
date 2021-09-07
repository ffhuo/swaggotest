package swaggo

import (
	"regexp"
	"strings"

	"github.com/ffhuo/swaggotest/internal/input"
	"github.com/ffhuo/swaggotest/internal/models"
	"github.com/ffhuo/swaggotest/internal/output"
	"github.com/pkg/errors"
)

var mustRegexp = regexp.MustCompile(`#\/definitions\/(.[a-zA-Z.]+)`)

type Options struct {
	Path string
}

// A GeneratedTest contains information about a test file with generated tests.
type GeneratedTest struct {
	Path      string             // The test file's absolute path.
	Functions []*models.Function // The functions with new test methods.
	Output    []byte             // The contents of the test file.
}

func GenerateTests(opt *Options) ([]*GeneratedTest, error) {
	swaggerData, err := input.ReadFile(opt.Path)
	if err != nil {
		return nil, errors.Errorf("read file %s err: %v", opt.Path, err)
	}
	return generateTests(swaggerData, opt)
}

func generateTests(swaggerData *models.SwaggerData, opt *Options) ([]*GeneratedTest, error) {
	var tests []*GeneratedTest
	tests = make([]*GeneratedTest, 0)

	var outopt *output.Options = &output.Options{
		PrintInputs: true,
		Subtests:    false,
		Parallel:    true,
		Named:       true,
	}

	funcs := make([]*models.Function, 0)
	for path, method := range swaggerData.Paths {
		nameArr := strings.Split(path, "/")
		name := nameArr[len(nameArr)-1]
		for way, data := range method {
			if data.Tags[0] != "account-discount" {
				continue
			}
			methodFunc := &models.Function{}
			methodFunc.Method = way
			methodFunc.Name = name + "_" + way
			methodFunc.Consumes = data.Consumes
			methodFunc.ReturnsError = true
			methodFunc.IsExported = false
			methodFunc.Parameters = generateFuncParams(data.Parameters)
			funcs = append(funcs, methodFunc)
		}
	}

	head := &models.Header{
		Package: "docs",
	}
	byt, err := outopt.Process(head, funcs)
	if err != nil {
		return nil, err
	}

	tests = append(tests, &GeneratedTest{
		Path:      opt.Path,
		Functions: funcs,
		Output:    byt,
	})

	return tests, nil
}

//
func generateFuncParams(params []models.SwaggerApiMethodParameter) []*models.Field {
	var fields []*models.Field
	fields = make([]*models.Field, 0)
	for _, param := range params {
		field := &models.Field{}
		field.Name = param.Name
		field.In = param.In
		field.Type = &models.Expression{}
		if param.Schema != nil {
			params := mustRegexp.FindStringSubmatch(param.Schema.Ref)
			if len(params) < 1 {
				continue
			}
			field.Type.Value = params[1]
		} else {
			field.Type.Value = getType(param.Type)
		}
		fields = append(fields, field)
	}
	return fields
}

// func findDefinitions(definitions map[string]models.SwaggerDefinitionData, def string) *models.Expression {
// 	_, ok := definitions[def]
// 	if !ok {
// 		return nil
// 	}

// 	var field *models.Field = &models.Field{}
// 	field.Type = &models.Expression{}
// 	field.Type.Value = def
// 	return field
// }

//
func getType(t string) string {
	switch t {
	case "boolean":
		return "bool"
	case "integer":
		return "int"
	case "long":
		return "int64"
	case "float":
		return "float32"
	case "double":
		return "float64"
	}
	return t
}
