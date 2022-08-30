package dz

import (
	"github.com/DolkMd/go-dezero/domain/core"
	"github.com/fcfcqloow/go-advance/log"
)

type Variables []Variable

func NewVariables(data ...core.Matrix) Variables {
	result := make([]Variable, 0, len(data))
	for _, v := range data {
		result = append(result, NewVariable(v))
	}

	return result
}

func (v Variables) First() Variable {
	if len(v) == 0 {
		return nil
	}

	return v[0]
}

func (v Variables) F() Variable {
	return v.First()
}

func (v Variables) Grads() Variables {
	result := []Variable{}
	for _, variable := range v {
		if variable.Grad() != nil {
			result = append(result, variable.Grad())
		} else {
			log.Warn("no grad")
		}
	}
	return result
}

func (v Variables) Generations() []int {
	result := []int{}
	for _, variable := range v {
		if variable != nil {
			result = append(result, variable.Generation())
		}
	}

	return result
}

func (v Variables) DataArr() []core.Matrix {
	result := make([]core.Matrix, 0, len(v))
	for _, v := range v {
		result = append(result, v.Data())
	}

	return result
}

func (v Variables) Shapes() []core.Shape {
	result := []core.Shape{}
	for _, v := range v {
		result = append(result, v.Shape())
	}

	return result
}

func (v Variables) Names() []string {
	result := []string{}
	for _, v := range v {
		result = append(result, v.Name()+",")
	}

	return result
}

func (v Variables) ToDataList() []core.Matrix {
	result := make([]core.Matrix, 0, len(v))
	for _, v := range v {
		result = append(result, v.Data())
	}

	return result
}
