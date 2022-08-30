package fn

import (
	"fmt"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type sum struct {
	dz.Function
	xShape  core.Shape
	options []core.Option
}

func NewSum(options ...core.Option) dz.Function {
	instance := new(sum)
	instance.Function = dz.ExtendsFunction(
		instance.Forward,
		instance.Backward,
		"Sum",
	)
	instance.options = options
	return instance
}

func (s *sum) Forward(variables ...dz.Variable) dz.Variables {
	x := variables[0]
	v := x.Data().Sum(s.options...)
	s.xShape = x.Data().Shape()
	return []dz.Variable{dz.NewVariable(v, dz.VarOpts().Name(fmt.Sprintf("(SUM(%s))", x.Name())))}
}

func (s *sum) Backward(variables ...dz.Variable) dz.Variables {
	gy := variables[0]
	gx := gy.Data().Broadcast(s.xShape)
	return dz.NewVariables(gx)
}
