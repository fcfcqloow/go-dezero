package fn_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestDiv(t *testing.T) {

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		"success: 10/2": {
			calc: func() []core.Matrix {
				a := dz.NewVariable(core.New1D(10))
				b := dz.NewVariable(core.New1D(2))
				return []core.Matrix{fn.Div(a, b).Data()}
			},
			result: dz.NewVariables(core.New1D(5)).DataArr(),
		},
		"success: 10/-2": {
			calc: func() []core.Matrix {
				a := dz.NewVariable(core.New1D(10))
				b := dz.NewVariable(core.New1D(-2))
				return []core.Matrix{fn.Div(a, b).Data()}
			},
			result: dz.NewVariables(core.New1D(-5)).DataArr(),
		},
		"success: backward": {
			calc: func() []core.Matrix {
				a := dz.NewVariable(core.New1D(1))
				b := dz.NewVariable(core.New1D(2))
				y := fn.Div(a, b)
				y.Backward()
				return []core.Matrix{a.Grad().Data(), b.Grad().Data()}
			},
			result: []core.Matrix{core.New1D(0.5), core.New1D(-0.25)},
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
