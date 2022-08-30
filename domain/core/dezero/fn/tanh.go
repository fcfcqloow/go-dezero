package fn

import (
	"math"

	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type (
	tanh struct{ dz.Function }
)

func NewTanh() dz.Function {
	instance := new(tanh)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "Tanh")
	return instance
}

func (*tanh) Forward(variables ...dz.Variable) dz.Variables {
	x := variables[0]
	return dz.NewVariables(x.Data().CopyApply(math.Tanh))
}

func (t *tanh) Backward(variables ...dz.Variable) dz.Variables {
	y := t.Outputs()[0]
	a := Mul(y, y)
	b := Sub(dz.NewVariable(a.Data().CopyFull(1)), a)
	return []dz.Variable{Mul(variables[0], b)}
}
