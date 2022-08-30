package fn

import (
	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type (
	sub struct {
		dz.Function
		x0Shape core.Shape
		x1Shape core.Shape
	}
)

func NewSub() dz.Function {
	instance := new(sub)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "Sub")
	return instance
}

func (s *sub) Forward(variables ...dz.Variable) dz.Variables {
	x0, x1 := variables[0], variables[1]
	s.x0Shape, s.x1Shape = x0.Shape(), x1.Shape()
	x0d, x1d := core.BroadcastForMatrix(x0.Data(), x1.Data())
	x0, x1 = dz.NewVariable(x0d), dz.NewVariable(x1d)
	return []dz.Variable{dz.NewVariable(x0.Data().CopySub(x1.Data()))}
}

func (s *sub) Backward(variables ...dz.Variable) dz.Variables {
	gx0 := variables[0]
	gx1 := NewNeg().Apply(gx0).First()
	if s.x0Shape != s.x1Shape {
		gx0 = SumTo(gx0, s.x0Shape)
		gx1 = SumTo(gx1, s.x1Shape)
	}
	return []dz.Variable{gx0, gx1}
}
