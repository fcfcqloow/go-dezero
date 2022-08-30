package fn

import (
	"math"

	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type (
	pow struct {
		dz.Function
		pow float64
	}
)

func NewPow(p float64) dz.Function {
	instance := new(pow)
	instance.pow = p
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "Pow")
	return instance
}

func (p *pow) Forward(variables ...dz.Variable) dz.Variables {
	y := variables[0].Data().CopyApply(func(v float64) float64 {
		return math.Pow(v, p.pow)
	})
	return []dz.Variable{dz.NewVariable(y)}
}

func (p *pow) Backward(variables ...dz.Variable) dz.Variables {
	gx := variables[0]
	x := p.Inputs()[0]
	c := p.pow
	a := NewPow(c - 1).Apply(x).First()
	b := Mul(a, gx)
	gy := Mul(b, dz.NewVariable(b.Data().CopyFull(c)))
	return []dz.Variable{gy} //c * x ** (c - 1) * gy
}
