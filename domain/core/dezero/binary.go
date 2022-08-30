package dz

import (
	"github.com/DolkMd/go-dezero/domain/core"
)

func SaveVariables(filename string, mp map[string]Variable) error {
	tmp := map[string]core.Matrix{}
	for k, v := range mp {
		tmp[k] = v.Data()
	}

	return core.SaveMatrix(filename, tmp)
}

func LoadVariables(filename string) (map[string]Variable, error) {
	fileValue, err := core.LoadMatrix(filename)
	if err != nil {
		return nil, err
	}
	result := map[string]Variable{}
	for k, v := range fileValue {
		result[k] = NewVariable(v)
	}

	return result, nil
}
