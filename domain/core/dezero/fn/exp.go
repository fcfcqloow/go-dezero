package fn

import (
	"math"

	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type exp struct{ dz.Function }

func NewExp() dz.Function {
	instance := new(exp)
	instance.Function = dz.ExtendsFunction(
		instance.Forward,
		instance.Backward,
		"Exp",
	)
	return instance
}

func (*exp) Forward(variables ...dz.Variable) dz.Variables {
	y := variables[0].Data().CopyApply(math.Exp)
	return []dz.Variable{dz.NewVariable(y)}
}

func (e *exp) Backward(variables ...dz.Variable) dz.Variables {
	x := e.Inputs()[0].Data()
	x.Apply(math.Exp)
	x = x.CopyMul(variables[0].Data())

	return dz.NewVariables(x)
}
