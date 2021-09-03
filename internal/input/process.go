package input

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ffhuo/swaggotest/internal/models"
	"github.com/pkg/errors"
)

func ReadFile(filePath string) (*models.SwaggerData, error) {
	var swaggerData models.SwaggerData
	byteData, err := readFile(filePath)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(byteData, &swaggerData); err != nil {
		return nil, errors.Wrap(err, "unmarshal swagger data err")
	}
	return &swaggerData, nil
}

func readFile(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "open: %s", filePath)
	}
	defer f.Close()

	byteData, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrapf(err, "read: %s", filePath)
	}

	return byteData, nil
}
