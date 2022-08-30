package fn

import (
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type (
	linear struct {
		dz.Function
	}
)

func NewLinear() dz.Function {
	instance := new(linear)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "Linear")
	return instance
}

func (l *linear) Forward(variables ...dz.Variable) dz.Variables {
	x, w, b := variables[0], variables[1], variables[2]
	y := MatMul(x, w)
	if b != nil {
		y = Add(y, b)
	}
	return []dz.Variable{y}
}

func (l *linear) Backward(variables ...dz.Variable) dz.Variables {
	gy := variables[0]
	x, W := l.Inputs()[0], l.Inputs()[1]

	var b dz.Variable
	if len(l.Inputs()) > 2 {
		b = l.Inputs()[2]
	}

	var gb dz.Variable
	if b == nil || b.Data() == nil {
		gb = nil
	} else {
		gb = SumTo(gy, b.Shape())
	}
	var gx, gW, tW, tx dz.Variable

	tx = Transpose(x)
	tW = Transpose(W)
	gx = MatMul(gy, tW)
	gW = MatMul(tx, gy)

	return []dz.Variable{gx, gW, gb}
}
