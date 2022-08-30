package fn

import (
	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type sigmoid struct{ dz.Function }

func NewSigmoid() dz.Function {
	instance := new(sigmoid)
	instance.Function = dz.ExtendsFunction(
		instance.Forward,
		instance.Backward,
		"Sigmoid",
	)
	return instance
}

func (*sigmoid) Forward(variables ...dz.Variable) dz.Variables {
	x := variables[0]
	a1 := Add(Exp(Neg(x)), dz.NewVariable(core.New1D(1)))
	a2 := dz.NewVariable(core.New1D(1))
	return []dz.Variable{Div(a2, a1)}
}

func (s *sigmoid) Backward(variables ...dz.Variable) dz.Variables {
	y, gy := s.Outputs()[0], variables[0]
	a1 := Mul(gy, y)
	gx := Mul(a1, Sub(dz.NewVariable(core.New1D(1)), y))
	return []dz.Variable{gx}
}
