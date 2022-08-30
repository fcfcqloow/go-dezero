package fn_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestSoftmaxCrossEntropy(t *testing.T) {

	t.Run("gradientCheck 1", func(t *testing.T) {

		x := dz.NewVariable(core.New2D([][]float64{
			{-1, 0, 1, 2},
			{2, 0, 1, -1},
		}))
		label := dz.NewVariable(core.New1D(3, 0))
		f := func(x dz.Variable) dz.Variable { return fn.SoftmaxCrossEntropy(x, label) }
		assert.True(t, gradientCheck(t, f, x))
	})

	// t.Run("gradientCheck 2", func(t *testing.T) {
	//
	// 	x0 := dz.NewVariable(core.NewRandN(core.Shape{1, 100}))
	// 	x1 := dz.NewVariable(core.NewRandN(core.Shape{1, 100}))
	// 	f := func(x dz.Variable) dz.Variable {
	// 		return fn.MeanSquaredError(x0, x1)
	// 	}
	// 	assert.True(t, gradientCheck(t, f, x0))
	// })
}
