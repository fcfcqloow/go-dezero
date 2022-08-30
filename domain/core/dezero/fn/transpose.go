package fn

import dz "github.com/DolkMd/go-dezero/domain/core/dezero"

type (
	transpose struct {
		dz.Function
	}
)

func NewTranspose() dz.Function {
	instance := new(transpose)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "Transpose")
	return instance
}

func (r *transpose) Forward(variables ...dz.Variable) dz.Variables {
	x := variables[0]
	y := x.Data().CopyT()
	return dz.NewVariables(y)
}

func (r *transpose) Backward(variables ...dz.Variable) dz.Variables {
	gy := variables[0]
	gx := NewTranspose().Apply(gy)
	return gx
}
