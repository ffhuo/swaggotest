package swaggo

import (
	"billing/common"

	"github.com/pkg/errors"
)

var apiMethods *SwaggerData

func ReadDoc(fileName string) (*SwaggerData, error) {
	if apiMethods != nil {
		return apiMethods, nil
	}

	var data SwaggerData
	if err := common.ReadFile(fileName, &data); err != nil {
		return nil, errors.Wrap(err, "read file err")
	}
	apiMethods = &data
	return apiMethods, nil
}

func ls(a, b int) int {
	return a + b
}
