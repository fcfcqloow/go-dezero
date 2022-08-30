package fn_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestMatMul(t *testing.T) {

	x := dz.NewVariable(core.New2D([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}))
	W := dz.NewVariable(core.New2D([][]float64{
		{1, 2, 3, 4},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}))
	y := fn.MatMul(x, W)
	y.Backward()

	assert.Equal(t, x.Grad().Shape(), core.Shape{R: 2, C: 3})
	assert.Equal(t, W.Grad().Shape(), core.Shape{R: 3, C: 4})

	t.Run("gradientCheck 1", func(t *testing.T) {
		x := dz.NewVariable(core.NewRandN(core.Shape{3, 2}))
		w := dz.NewVariable(core.NewRandN(core.Shape{2, 3}))
		f := func(dz.Variable) dz.Variable {
			return fn.MatMul(x, w)
		}
		assert.True(t, gradientCheck(t, f, x))
	})

	t.Run("gradientCheck 2", func(t *testing.T) {
		x := dz.NewVariable(core.NewRandN(core.Shape{5, 2}))
		w := dz.NewVariable(core.NewRandN(core.Shape{2, 9}))
		f := func(dz.Variable) dz.Variable {
			return fn.MatMul(x, w)
		}
		assert.True(t, gradientCheck(t, f, x))
	})
}
