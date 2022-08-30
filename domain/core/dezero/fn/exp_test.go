package fn_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestExp(t *testing.T) {

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		"success: exp e^2.1=8.16616991256765": {
			calc: func() []core.Matrix {
				a := dz.NewVariable(core.New1D(2.1))
				return []core.Matrix{fn.Exp(a).Data()}
			},
			result: []core.Matrix{core.New1D(8.16616991256765)},
		},
		"success: backward exp(x^2)^2": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New1D(0.5))
				y := fn.Square(x)
				y = fn.Exp(y)
				y = fn.Square(y)
				y.SetGrad(dz.NewVariable(core.New1D(1.0)))
				y.Backward()
				return []core.Matrix{y.Data(), x.Grad().Data()}
			},
			result: []core.Matrix{core.New1D(1.648721270700128), core.New1D(3.297442541400256)},
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
