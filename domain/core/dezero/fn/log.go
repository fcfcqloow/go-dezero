package fn

import dz "github.com/DolkMd/go-dezero/domain/core/dezero"

type (
	_log struct {
		dz.Function
	}
)

func NewLog() dz.Function {
	instance := new(_log)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "Log")
	return instance
}

func (l *_log) Forward(variables ...dz.Variable) dz.Variables {
	x := variables[0]
	y := dz.NewVariable(x.Data().CopyLog())
	return []dz.Variable{y}
}

func (l *_log) Backward(variables ...dz.Variable) dz.Variables {
	x := l.Inputs()[0]
	gx := Div(variables[0], x)
	return []dz.Variable{gx}
}
