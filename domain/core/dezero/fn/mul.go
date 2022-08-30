package fn

import (
	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type (
	mul struct{ dz.Function }
)

func NewMul() dz.Function {
	instance := new(mul)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "Mul")
	return instance
}

func (m *mul) Forward(variables ...dz.Variable) dz.Variables {
	x0, x1 := variables[0], variables[1]
	x0d, x1d := core.BroadcastForMatrix(x0.Data(), x1.Data())
	x0, x1 = dz.NewVariable(x0d), dz.NewVariable(x1d)
	return []dz.Variable{dz.NewVariable(x0.Data().CopyMul(x1.Data()))}
}

func (m *mul) Backward(variables ...dz.Variable) dz.Variables {
	x0, x1 := m.Inputs()[0], m.Inputs()[1]
	gx0 := Mul(variables[0], x1)
	gx1 := Mul(variables[0], x0)
	if x0.Shape() != x1.Shape() {
		gx0 = SumTo(gx0, x0.Shape())
		gx1 = SumTo(gx1, x1.Shape())
	}
	return []dz.Variable{gx0, gx1}
}
