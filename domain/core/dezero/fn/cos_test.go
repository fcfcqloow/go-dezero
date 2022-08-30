package fn_test

import (
	"math"
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestCos(t *testing.T) {

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		"success: backward": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New1D(math.Pi / 4.0))
				y := fn.Cos(x)
				y.Backward()
				return []core.Matrix{x.Grad().Data(), y.Data()}
			},
			result: []core.Matrix{core.New1D(-0.7071067811865475), core.New1D(0.7071067811865476)},
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
