package fn

import (
	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type (
	relu struct {
		dz.Function
		x0Shape core.Shape
		x1Shape core.Shape
	}
)

func NewRelu() dz.Function {
	instance := new(relu)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "Relu")
	return instance
}

func (r *relu) Forward(variables ...dz.Variable) dz.Variables {
	x := variables[0]
	y := x.Data().CopyApply(func(f float64) float64 {
		if f > 0 {
			return f
		}
		return 0
	})

	return []dz.Variable{dz.AsVariable(y)}
}

func (r *relu) Backward(variables ...dz.Variable) dz.Variables {
	x := r.Inputs()[0]
	mask := x.Data().OnOff(func(i, j int, v float64) bool {
		return v > 0
	})
	gx := Mul(variables[0], dz.AsVariable(mask))
	return []dz.Variable{gx}

}
