package fn

import (
	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type (
	bt struct {
		dz.Function
		shape  core.Shape
		xShape core.Shape
	}

	add struct {
		dz.Function
		x0Shape core.Shape
		x1Shape core.Shape
	}
)

func NewAdd() dz.Function {
	instance := new(add)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "Add")
	return instance
}

func (a *add) Forward(variables ...dz.Variable) dz.Variables {
	x0, x1 := variables[0], variables[1]
	a.x0Shape, a.x1Shape = x0.Shape(), x1.Shape()
	x0d, x1d := core.BroadcastForMatrix(x0.Data(), x1.Data())
	x0, x1 = dz.NewVariable(x0d, dz.VarOpts().Name(x0.Name())), dz.NewVariable(x1d, dz.VarOpts().Name(x1.Name()))
	return []dz.Variable{dz.NewVariable(x0.Data().CopyAdd(x1.Data()))}
}

func (a *add) Backward(variables ...dz.Variable) dz.Variables {
	gx0, gx1 := variables[0], variables[0]
	if a.x0Shape != a.x1Shape {
		gx0 = NewSumTo(a.x0Shape).Apply(gx0)[0]
		gx1 = NewSumTo(a.x1Shape).Apply(gx0)[0]
	}

	return []dz.Variable{gx0, gx1}

}

func NewBroadcastTo(shape core.Shape) dz.Function {
	instance := new(bt)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "BroadcastTo")
	instance.shape = shape
	return instance
}

func (b *bt) Forward(variables ...dz.Variable) dz.Variables {
	x := variables[0]
	b.xShape = x.Data().Shape()
	return dz.NewVariables(x.Data().Broadcast(b.shape))
}

func (b *bt) Backward(variables ...dz.Variable) dz.Variables {
	return NewSumTo(b.xShape).Apply(variables[0])
}
