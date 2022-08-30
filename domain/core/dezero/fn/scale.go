package fn

import dz "github.com/DolkMd/go-dezero/domain/core/dezero"

type (
	scale struct {
		dz.Function
		v float64
	}
)

func NewScale(v float64) dz.Function {
	instance := new(scale)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "Scale")
	instance.v = v
	return instance
}

func (s *scale) Forward(variables ...dz.Variable) dz.Variables {
	x := variables[0]
	return dz.NewVariables(x.Data().CopyScale(s.v))
}

func (s *scale) Backward(variables ...dz.Variable) dz.Variables {
	x := s.Inputs()[0]
	return NewScale(s.v).Apply(x)
}
