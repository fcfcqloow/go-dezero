package dz

import (
	"fmt"

	"github.com/DolkMd/go-dezero/domain/core"

	xmath "github.com/fcfcqloow/go-advance/math"
)

type (
	Forward  func(...Variable) Variables
	Backward func(...Variable) Variables
	Function interface {
		fmt.Stringer
		Apply(...Variable) Variables
		Inputs() Variables
		SetInputs(vs Variables)
		Outputs() Variables
		Backward(...Variable) Variables
		Forward(...Variable) Variables
		Generation() int
		Name() string
	}
	function struct {
		forward    Forward
		backward   Backward
		inputs     Variables
		outputs    Variables
		generation int
		name       string
	}
)

func ExtendsFunction(forward Forward, backward Backward, name string) Function {
	return &function{
		forward:  forward,
		backward: backward,
		name:     name,
	}
}

func (f *function) Apply(inputs ...Variable) Variables {
	xs := Variables(inputs)
	outputs := f.forward(xs...)
	if core.Config.EnableBackprop() {
		f.generation = xmath.MaxInt(Variables(inputs).Generations()...)

		for i := range outputs {
			outputs[i].SetCreator(f)
		}
		for _, input := range inputs {
			if input != nil {
				f.inputs = append(f.inputs, input)
			}
		}
		f.outputs = outputs

	}
	// log.Debug(gputil.Ellipsis(fmt.Sprintf("Fw: %s, G: %d", f.String(), f.Generation()), 50))

	return outputs
}
func (f *function) Inputs() Variables {
	return f.inputs
}

func (f *function) SetInputs(vs Variables) {
	f.inputs = vs
}

func (f *function) Outputs() Variables {
	return f.outputs
}

func (f *function) Backward(dence ...Variable) Variables {
	return f.backward(dence...)
}

func (f *function) Forward(dence ...Variable) Variables {
	return f.forward(dence...)
}

func (f *function) Generation() int {
	return f.generation
}

func (f *function) Name() string {
	return f.name
}

func (f *function) String() string {
	return fmt.Sprintf("%s: %p", f.name, f)
}

func NumericalDiff(f func(Variable) Variable, x Variable, eps *float64) core.Matrix {
	if eps == nil {
		eps = new(float64)
		*eps = 1e-4
	}

	x0 := NewVariable(x.Data().CopyApply(func(v float64) float64 { return v - *eps }))
	x1 := NewVariable(x.Data().CopyApply(func(v float64) float64 { return v + *eps }))
	y0 := f(x0)
	y1 := f(x1)

	result := y1.Data().CopySub(y0.Data())
	result.Apply(func(v float64) float64 { return v / (2 * (*eps)) })
	return result
}
