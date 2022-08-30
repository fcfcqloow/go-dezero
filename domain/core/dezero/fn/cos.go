package fn

import (
	"math"

	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type cos struct{ dz.Function }

func NewCos() dz.Function {
	instance := new(cos)
	instance.Function = dz.ExtendsFunction(
		instance.Forward,
		instance.Backward,
		"Cos",
	)
	return instance
}

func (*cos) Forward(variables ...dz.Variable) dz.Variables {
	x := variables[0]
	y := x.Data().CopyApply(func(f float64) float64 { return math.Cos(f) })
	return dz.NewVariables(y)
}

func (e *cos) Backward(variables ...dz.Variable) dz.Variables {
	x := e.Inputs()[0]
	gy := variables[0]
	gx := Mul(gy, Neg(Sin(x)))
	return []dz.Variable{gx}
}
