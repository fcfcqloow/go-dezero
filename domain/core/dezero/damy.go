package dz

import (
	"fmt"

	"github.com/DolkMd/go-dezero/domain/core"
)

func _add(x0, x1 Variable) Variable { return NewAdd().Apply(x0, x1).First() }

type (
	bt struct {
		Function
		shape  core.Shape
		xShape core.Shape
	}
	sum struct {
		Function
		xShape  core.Shape
		options []core.Option
	}
	sumTo struct {
		Function
		xShape  core.Shape
		shape   core.Shape
		options []core.Option
	}
	add struct {
		Function
		x0Shape core.Shape
		x1Shape core.Shape
	}
)

func NewAdd() Function {
	instance := new(add)
	instance.Function = ExtendsFunction(instance.Forward, instance.Backward, "Add")
	return instance
}

func (a *add) Forward(variables ...Variable) Variables {
	x0, x1 := variables[0], variables[1]
	a.x0Shape, a.x1Shape = x0.Shape(), x1.Shape()
	x0d, x1d := core.BroadcastForMatrix(x0.Data(), x1.Data())
	x0, x1 = NewVariable(x0d, VarOpts().Name(x0.Name())), NewVariable(x1d, VarOpts().Name(x1.Name()))
	return []Variable{NewVariable(x0.Data().CopyAdd(x1.Data()))}
}

func (a *add) Backward(variables ...Variable) Variables {
	gx0, gx1 := variables[0], variables[0]
	if a.x0Shape != a.x1Shape {
		gx0 = NewSumTo(a.x0Shape).Apply(gx0)[0]
		gx1 = NewSumTo(a.x1Shape).Apply(gx0)[0]
	}

	return []Variable{gx0, gx1}

}

func NewBroadcastTo(shape core.Shape) Function {
	instance := new(bt)
	instance.Function = ExtendsFunction(instance.Forward, instance.Backward, "BroadcastTo")
	instance.shape = shape
	return instance
}

func (b *bt) Forward(variables ...Variable) Variables {
	x := variables[0]
	b.xShape = x.Data().Shape()
	return NewVariables(x.Data().Broadcast(b.shape))
}

func (b *bt) Backward(variables ...Variable) Variables {
	return NewSumTo(b.xShape).Apply(variables[0])
}

func NewSumTo(shape core.Shape, options ...core.Option) Function {
	instance := new(sumTo)
	instance.Function = ExtendsFunction(instance.Forward, instance.Backward, "SumTo")
	instance.shape = shape
	instance.options = options
	return instance
}

func (s *sumTo) Forward(variables ...Variable) Variables {
	x := variables[0]
	xshape := x.Shape()
	if s.shape == xshape {
		return []Variable{x}
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

func (s *sumTo) Backward(variables ...Variable) Variables {
	return NewBroadcastTo(s.xShape).Apply(variables[0])
}

func NewSum(options ...core.Option) Function {
	instance := new(sum)
	instance.Function = ExtendsFunction(
		instance.Forward,
		instance.Backward,
		"Sum",
	)
	instance.options = options
	return instance
}

func (s *sum) Forward(variables ...Variable) Variables {
	x := variables[0]
	v := x.Data().Sum(s.options...)
	s.xShape = x.Data().Shape()
	return []Variable{NewVariable(v, VarOpts().Name(fmt.Sprintf("(SUM(%s))", x.Name())))}
}

func (s *sum) Backward(variables ...Variable) Variables {
	gy := variables[0]
	gx := gy.Data().Broadcast(s.xShape)
	return NewVariables(gx)
}
