package fn_test

import (
	"fmt"
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		"success: test Square model": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New1D(2.0))
				y := fn.Square(x)
				return []core.Matrix{y.Data()}
			},
			result: []core.Matrix{core.New1D(4.0)},
		},
		"success: test Square model backward": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New1D(3.0))
				y := fn.Square(x)
				y.Backward()
				return []core.Matrix{x.Grad().Data()}
			},
			result: []core.Matrix{core.New1D(6.0)},
		},
		"x^2^2 + x^2^2 = 2(x^4)= 32": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New1D(2.0))    // x = 2
				a := fn.Square(x)                       // x^2 = 4
				y := fn.Add(fn.Square(a), fn.Square(a)) // x^2^2 + x^2^2 = 2(x^4)= 32
				y.Backward()                            // y' = 8(x^3) = 64

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

	t.Run("success: Square model gradient", func(t *testing.T) {
		x := dz.NewVariable(core.New1D(random(t, 0, 1)))
		y := fn.Square(x)
		y.Backward()

		gard := dz.NumericalDiff(fn.Square, x, nil)
		flag := x.Grad().Data().Allclose(gard, nil, nil)
		t.Log(fmt.Sprintf("grad: %v, x-grad: %v", gard.At(0, 0), x.Grad().Data().At(0, 0)))
		assert.True(t, flag)
	})
}
