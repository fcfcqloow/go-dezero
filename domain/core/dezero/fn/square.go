package fn

import dz "github.com/DolkMd/go-dezero/domain/core/dezero"

type square struct {
	dz.Function
}

func NewSquare() dz.Function {
	instance := new(square)
	instance.Function = dz.ExtendsFunction(
		instance.Forward,
		instance.Backward,
		"Square",
	)
	return instance
}

func (*square) Forward(variables ...dz.Variable) dz.Variables {
	return NewPow(2).Forward(variables[0])
}

func (s *square) Backward(variables ...dz.Variable) dz.Variables {
	y := Mul(s.Inputs()[0], dz.NewVariable(s.Inputs().F().Data().CopyFull(2)))
	gy := Mul(y, variables[0])
	return []dz.Variable{gy}
}
