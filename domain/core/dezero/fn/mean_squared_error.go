package fn

import (
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type (
	meanSquaredError struct {
		dz.Function
	}
)

var cnt = 0

func NewMeanSquaredError() dz.Function {
	instance := new(meanSquaredError)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "MeanSquaredError")
	return instance
}

func (*meanSquaredError) Forward(variables ...dz.Variable) dz.Variables {
	x0, x1 := variables[0], variables[1]
	diff := Sub(x0, x1)
	temp := Sum(Pow(diff, 2))
	lenDiff := diff.Shape().R
	tempd := temp.Data().CopyFull(float64(lenDiff))
	return []dz.Variable{Div(temp, dz.NewVariable(tempd))}
}

func (m *meanSquaredError) Backward(variables ...dz.Variable) dz.Variables {
	x0, x1, gy := m.Inputs()[0], m.Inputs()[1], variables[0]
	diff := Sub(x0, x1)
	gy = NewBroadcastTo(diff.Shape()).Apply(gy)[0]
	a1 := 2.0 / float64(diff.Shape().R)
	a2 := Mul(gy, diff)
	gx0 := Mul(a2, dz.NewVariable(a2.Data().CopyFull(float64(a1))))
	gx1 := Neg(gx0)
	return []dz.Variable{gx0, gx1}
}
