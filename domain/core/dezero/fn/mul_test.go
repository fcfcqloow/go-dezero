package fn_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestMul(t *testing.T) {

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		"success: a*b + c": {
			calc: func() []core.Matrix {
				a := dz.NewVariable(core.New1D(3.0))
				b := dz.NewVariable(core.New1D(2.0))
				c := dz.NewVariable(core.New1D(1.0))
				y := fn.Add(fn.Mul(a, b), c)
				y.Backward()
				return []core.Matrix{y.Data(), a.Grad().Data(), b.Grad().Data()}
			},
			result: []core.Matrix{core.New1D(7), core.New1D(2), core.New1D(3)},
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
