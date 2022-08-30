package fn

import (
	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type (
	sumTo struct {
		dz.Function
		xShape  core.Shape
		shape   core.Shape
		options []core.Option
	}
)

func NewSumTo(shape core.Shape, options ...core.Option) dz.Function {
	instance := new(sumTo)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "SumTo")
	instance.shape = shape
	instance.options = options
	return instance
}

func (s *sumTo) Forward(variables ...dz.Variable) dz.Variables {
	x := variables[0]
	xshape := x.Shape()
	if s.shape == xshape {
		return []dz.Variable{x}
	}
	if s.shape.R == 1 && s.shape.C == 1 {
		return NewSum().Apply(x)
	}
	if s.shape.R == 1 {
		return NewSum(core.Axis(0)).Apply(x)
	}

	if s.shape.R > 1 {
		return NewSum(core.Axis(1)).Apply(x)
	}

	panic("sum to Forward")
}

func (s *sumTo) Backward(variables ...dz.Variable) dz.Variables {
	return NewBroadcastTo(s.xShape).Apply(variables[0])
}
