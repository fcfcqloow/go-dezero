package fn

import (
	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type (
	reshape struct {
		dz.Function
		shape, xshape core.Shape
	}
)

func NewReshape(shape core.Shape) dz.Function {
	instance := new(reshape)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "Reshape")
	instance.shape = shape
	return instance
}

func (r *reshape) Forward(variables ...dz.Variable) dz.Variables {
	x := variables[0]
	r.xshape = x.Data().Shape()
	y := x.Data().CopyReshape(r.shape)
	return dz.NewVariables(y)
}

func (r *reshape) Backward(variables ...dz.Variable) dz.Variables {
	gy := variables[0]
	return NewReshape(r.xshape).Apply(gy)
}
