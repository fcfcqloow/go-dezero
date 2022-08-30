package fn_test

import (
	"math"
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestSin(t *testing.T) {

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		"success": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New1D(1))
				return []core.Matrix{fn.Sin(x).Data()}
			},
			result: []core.Matrix{core.New1D(0.8414709848078965)},
		},
		"success: backward": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New1D(math.Pi / 4.0))
				y := fn.Sin(x)
				y.Backward()
				return []core.Matrix{x.Grad().Data(), y.Data()}
			},
			result: []core.Matrix{core.New1D(0.7071067811865476), core.New1D(0.7071067811865475)},
		},
		"success: CreateGraph(true)": {
			calc: func() (result []core.Matrix) {

				x := dz.NewVariable(core.New1D(1.0))
				y := fn.Sin(x)
				y.Backward(dz.CreateGraph(true))

				for i := 0; i < 3; i++ {
					gx := x.Grad()
					x.ClearGrad()
					gx.Backward(dz.CreateGraph(true))
					result = append(result, x.Grad().Data())
				}
				return
			},
			result: []core.Matrix{core.New1D(-0.8414709848078965), core.New1D(-0.5403023058681398), core.New1D(0.8414709848078965)},
		},
		"success: 2D": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New2D([][]float64{
					{1, 2, 3},
					{4, 5, 6},
				}))
				y := fn.Sin(x)
				return []core.Matrix{y.Data()}
			},
			result: []core.Matrix{core.New2D([][]float64{
				{0.8414709848078965, 0.9092974268256816, 0.1411200080598672},
				{-0.7568024953079282, -0.9589242746631385, -0.27941549819892586},
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
