package swaggo

import (
	"strings"

	"github.com/ffhuo/swaggotest/internal/input"
	"github.com/ffhuo/swaggotest/internal/models"
	"github.com/pkg/errors"
)

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
	for path, method := range swaggerData.Paths {
		nameArr := strings.Split(path, "/")
		name := nameArr[len(nameArr)]
		for way, data := range method {
			methodFunc := &models.Function{}
			methodFunc.Method = way
			methodFunc.Name = name + "_" + way
			methodFunc.Consumes = data.Consumes
			methodFunc.ReturnsError = false
			methodFunc.IsExported = false
		}
	}
	return tests, nil
}

//
func generateFuncParams(params []models.SwaggerApiMethodParameter) []*models.Field {
	var fields []*models.Field
	fields = make([]*models.Field, 0)
	for _, param := range params {
		filed := &models.Field{}
		filed.Name = param.Name
		filed.In = param.In
		filed.Type = &models.Expression{}
		if param.Schema != nil {
			// TODO 查找
		} else {
			filed.Type.Value = getType(param.Type)
		}
	}
	return fields
}

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
