package fn

import dz "github.com/DolkMd/go-dezero/domain/core/dezero"

type (
	neg struct{ dz.Function }
)

func NewNeg() dz.Function {
	instance := new(neg)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "Neg")
	return instance
}

func (m *neg) Forward(variables ...dz.Variable) dz.Variables {
	y := variables[0].Data().CopyApply(func(v float64) float64 {
		return v * -1
	})
	return []dz.Variable{dz.NewVariable(y)}
}

func (m *neg) Backward(variables ...dz.Variable) dz.Variables {
	return NewNeg().Apply(variables...)
}
