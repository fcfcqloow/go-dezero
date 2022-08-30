package fn_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestLinear(t *testing.T) {
	t.Run("allClose", func(t *testing.T) {
		x := dz.AsVariable([][]float64{
			{1, 2, 3}, {4, 5, 6},
		})
		w := dz.NewVariable(x.Data().CopyT())
		y := fn.Linear(x, w, nil)

		expexted := dz.AsVariable([][]float64{
			{14, 32}, {32, 77},
		})
		assert.True(t, allClose(t, expexted, y))
	})
	t.Run("gradientCheck", func(t *testing.T) {
		x := dz.AsVariable(core.NewRandN(core.Shape{3, 2}))
		w := dz.AsVariable(core.NewRandN(core.Shape{2, 3}))
		b := dz.AsVariable(core.NewRandN(core.Shape{1, 3}))
		f := func(x dz.Variable) dz.Variable {
			return fn.Linear(x, w, b)
		}
		assert.True(t, gradientCheck(t, f, x))
	})
}
