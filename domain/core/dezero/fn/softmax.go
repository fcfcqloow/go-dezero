package fn

import (
	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type (
	softmax struct {
		dz.Function
		options []core.Option
	}
)

func NewSoftmax(options ...core.Option) dz.Function {
	instance := new(softmax)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "Softmax")
	instance.options = options
	return instance
}

func (s *softmax) Forward(variables ...dz.Variable) dz.Variables {
	x := variables[0]
	a := x.Data().Max(s.options...)
	y := Sub(x, dz.NewVariable(a))
	y = Exp(y)
	y = Div(y, dz.NewVariable(y.Sum(s.options...)))
	return []dz.Variable{y}
}

func (s *softmax) Backward(variables ...dz.Variable) dz.Variables {
	gy := variables[0]
	y := s.Inputs()[0]
	gx := Mul(y, gy)
	sumdx := dz.NewVariable(gx.Sum(s.options...))
	gx = Sub(gx, Div(y, sumdx))
	return []dz.Variable{gx}
}
