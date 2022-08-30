package fn_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		"success: backward": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New1D(1, 2, 3, 4, 5, 6))
				y := fn.Sum(x)
				y.Backward()
				return []core.Matrix{y.Data(), x.Grad().Data()}
			},
			result: []core.Matrix{
				core.New1D(21),
				core.New1D(1, 1, 1, 1, 1, 1),
			},
		},
		"success: 2D": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New2D([][]float64{
					{1, 2, 3},
					{4, 5, 6},
				}))
				y := fn.Sum(x)
				y.Backward()

				return []core.Matrix{y.Data(), x.Grad().Data()}
			},
			result: []core.Matrix{
				core.New1D(21),
				core.New2D([][]float64{
					{1, 1, 1},
					{1, 1, 1},
				}),
			},
		},
		"success: Axis": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New2D([][]float64{
					{1, 2, 3},
					{4, 5, 6},
				}))
				y := fn.Sum(x, core.Axis(0))
				y.Backward()
				return []core.Matrix{y.Data(), x.Grad().Data()}
			},
			result: []core.Matrix{
				core.New1D(5, 7, 9),
				core.New2D([][]float64{
					{1, 1, 1},
					{1, 1, 1},
				}),
			},
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
