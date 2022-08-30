package dz

import "github.com/DolkMd/go-dezero/domain/core"

type (
	Parameter Variable
	parameter struct {
		Variable
	}
)

func NewParameter(m core.Matrix, opts ...VariableOpt) Parameter {
	return &parameter{Variable: NewVariable(m, opts...)}
}
func (p *parameter) String() string {
	return p.Name()
}

func IsParameter(v interface{}) bool {
	_, ok := v.(*parameter)
	return ok
}
