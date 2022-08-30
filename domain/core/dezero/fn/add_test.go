package fn_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		"success: add 2.1+2.8 = 4.9": {
			calc: func() []core.Matrix {
				a := dz.NewVariable(core.New1D(2.1))
				b := dz.NewVariable(core.New1D(2.8))
				return []core.Matrix{fn.Add(a, b).Data()}
			},
			result: dz.NewVariables(core.New1D(4.9)).DataArr(),
		},
		"success: add 3+2 = 5": {
			calc: func() []core.Matrix {
				x1, x2 := core.New1D(2), core.New1D(3)
				return []core.Matrix{fn.Add(dz.NewVariable(x1), dz.NewVariable(x2)).Data()}
			},
			result: dz.NewVariables(core.New1D(5)).DataArr(),
		},
		"success: broadcast": {
			calc: func() []core.Matrix {
				x0 := dz.NewVariable(core.New1D(1, 2, 3))
				x1 := dz.NewVariable(core.New1D(10))
				y := fn.Add(x0, x1)
				y.Backward()
				return []core.Matrix{y.Data(), x1.Grad().Data()}
			},
			result: dz.NewVariables(
				core.New2D([][]float64{{11, 12, 13}}),
				core.New1D(3),
			).DataArr(),
		},
		"success: 2D": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New2D([][]float64{{1, 2, 3}, {4, 5, 6}}))
				c := dz.NewVariable(core.New2D([][]float64{{10, 20, 30}, {40, 50, 60}}))
				y2 := fn.Add(x, c)
				return []core.Matrix{y2.Data()}
			},
			result: dz.NewVariables(
				core.New2D([][]float64{
					{11, 22, 33},
					{44, 55, 66},
				}),
			).DataArr(),
		},
		"success: backforward": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New1D(3.0))
				y := fn.Add(x, x)
				y.Backward()
				return []core.Matrix{x.Grad().Data()}
			},
			result: []core.Matrix{core.New1D(2)},
		},
		"success backforward2": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New1D(2.0))    // x = 2
				a := fn.Square(x)                       // x^2 = 4
				y := fn.Add(fn.Square(a), fn.Square(a)) // x^2^2 + x^2^2 = 2(x^4)= 32
				y.Backward()

				return []core.Matrix{y.Data(), x.Grad().Data()}
			},
			result: []core.Matrix{core.New1D(32), core.New1D(64)},
		},
	}

	for name, tc := range testCases {
		tc := tc
		name := name
		t.Run(name, func(t *testing.T) {

			assert.Equal(t, tc.result, tc.calc())
		})
	}
}
