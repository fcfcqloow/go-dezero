package fn

import (
	"math"

	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type sin struct {
	dz.Function
}

func NewSin() dz.Function {
	instance := new(sin)
	instance.Function = dz.ExtendsFunction(
		instance.Forward,
		instance.Backward,
		"Sin",
	)
	return instance
}

func (*sin) Forward(variables ...dz.Variable) dz.Variables {
	x := variables[0]
	y := x.Data().CopyApply(math.Sin)
	return dz.NewVariables(y)
}

func (e *sin) Backward(variables ...dz.Variable) dz.Variables {
	x := e.Inputs()[0]
	gy := variables[0]
	gx := Mul(gy, Cos(x))
	return []dz.Variable{gx}
}
