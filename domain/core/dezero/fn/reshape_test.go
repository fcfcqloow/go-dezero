package fn_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestReshape(t *testing.T) {

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		"success": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New2D([][]float64{
					{1, 2, 3},
					{4, 5, 6},
				}))
				x.Reshape(core.Shape{R: 1, C: 6})
				return []core.Matrix{x.Data()}
			},
			result: []core.Matrix{core.New1D(1, 2, 3, 4, 5, 6)},
		},
		"success: backward": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New2D([][]float64{
					{1, 2, 3},
					{4, 5, 6},
				}))
				y := fn.Reshape(x, core.Shape{R: 1, C: 6})
				y.Backward(dz.RetainGrad(true))
				return []core.Matrix{x.Grad().Data()}
			},
			result: []core.Matrix{core.New2D([][]float64{
				{1, 1, 1},
				{1, 1, 1},
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
