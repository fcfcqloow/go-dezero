package fn_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestSumTo(t *testing.T) {

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		"success: backward": {
			calc: func() []core.Matrix {
				result := []core.Matrix{}
				x := dz.NewVariable(core.New2D([][]float64{
					{1, 2, 3},
					{4, 5, 6},
				}))
				y := fn.SumTo(x, core.Shape{R: 1, C: 3})
				result = append(result, y.Data())

				y = fn.SumTo(x, core.Shape{R: 2, C: 1})
				result = append(result, y.Data())

				x0 := dz.NewVariable(core.New1D(1, 2, 3))
				x1 := dz.NewVariable(core.New1D(10))
				y = fn.Add(x0, x1)
				y.Backward()
				result = append(result, y.Data(), x1.Grad().Data())

				return result
			},
			result: []core.Matrix{
				core.New1D(5, 7, 9),
				core.New2D([][]float64{
					{6},
					{15},
				}),
				core.New1D(11, 12, 13),
				core.New1D(3),
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
