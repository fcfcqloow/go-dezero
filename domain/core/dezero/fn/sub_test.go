package fn_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		"success: 3-2": {
			calc: func() []core.Matrix {
				a := dz.NewVariable(core.New1D(3))
				b := dz.NewVariable(core.New1D(2))
				y := fn.Sub(a, b)
				y.Backward()
				return []core.Matrix{y.Data(), a.Grad().Data()}
			},
			result: []core.Matrix{core.New1D(1), core.New1D(1)},
		},
		"success: broadcast 1": {
			calc: func() []core.Matrix {
				a := dz.NewVariable(core.New2D([][]float64{
					{1, 2, 3, 4},
				}))
				b := dz.NewVariable(core.New1D(2))
				y := fn.Sub(a, b)
				y.Backward()
				return []core.Matrix{y.Data()}
			},
			result: []core.Matrix{core.New2D([][]float64{
				{-1, 0, 1, 2},
			})},
		},
		"success: broadcast 2": {
			calc: func() []core.Matrix {
				a := dz.NewVariable(core.New2D([][]float64{
					{1, 2, 3, 4},
					{1, 2, 3, 4},
					{1, 2, 3, 4},
				}))
				b := dz.NewVariable(core.New1D(2))
				y := fn.Sub(a, b)
				y.Backward()
				return []core.Matrix{y.Data()}
			},
			result: []core.Matrix{core.New2D([][]float64{
				{-1, 0, 1, 2},
				{-1, 0, 1, 2},
				{-1, 0, 1, 2},
			})},
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
