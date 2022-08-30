package fn_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestPow(t *testing.T) {

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		"success: x^3": {
			calc: func() []core.Matrix {
				a := dz.NewVariable(core.New1D(2.0))
				return []core.Matrix{fn.Pow(a, 3).Data()}
			},
			result: []core.Matrix{core.New1D(8)},
		},
		"success: CreateGraph": {
			calc: func() []core.Matrix {
				result := []core.Matrix{}
				f := func(x dz.Variable) dz.Variable {
					a := fn.Pow(x, 4)                 // x^4
					b := fn.MulFloat(fn.Pow(x, 2), 2) // x^2 * 2
					return fn.Sub(a, b)               // x^4 - 2(x^2)
				}
				x := dz.NewVariable(core.New1D(2.0))
				y := f(x)
				y.Backward(dz.CreateGraph(true)) // y' = 4x^3 - 4x =
				result = append(result, x.Grad().Data())

				gx := x.Grad()
				x.ClearGrad()
				gx.Backward() // y'' = 12x^2 - 4
				result = append(result, x.Grad().Data())
				return result
			},
			result: []core.Matrix{core.New1D(24), core.New1D(44)},
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
