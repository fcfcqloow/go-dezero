package fn

import (
	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type (
	div struct{ dz.Function }
)

func NewDiv() dz.Function {
	instance := new(div)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "Div")
	return instance
}

func (d *div) Forward(variables ...dz.Variable) dz.Variables {
	x0, x1 := variables[0], variables[1]
	x0d, x1d := core.BroadcastForMatrix(x0.Data(), x1.Data())
	x0, x1 = dz.NewVariable(x0d), dz.NewVariable(x1d)
	return []dz.Variable{dz.NewVariable(x0.Data().CopyDiv(x1.Data()))}
}

func (d *div) Backward(variables ...dz.Variable) dz.Variables {
	gy := variables[0]
	x0, x1 := d.Inputs()[0], d.Inputs()[1]
	gx0 := Div(gy, x1)
	x1_square := Square(x1)                        // x1^2
	x0_neg := Neg(x0)                              // x0 * -1
	x0_neg_div_x1_square := Div(x0_neg, x1_square) // -x0 / (x1^2)
	gx1 := Mul(gy, x0_neg_div_x1_square)           // gx * (-x0 / (x1^2))
	if x0.Shape() != x1.Shape() {
		gx0 = SumTo(gx0, x0.Shape())
		gx1 = SumTo(gx1, x1.Shape())
	}
	return []dz.Variable{gx0, gx1}
}
