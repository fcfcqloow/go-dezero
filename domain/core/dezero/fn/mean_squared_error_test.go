package fn_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestMeanSquaredError(t *testing.T) {
	t.Run("allclose 1", func(t *testing.T) {

		x0 := dz.NewVariable(core.New1D(0.0, 1.0, 2.0))
		x1 := dz.NewVariable(core.New1D(0.0, 1.0, 2.0))
		expected := fn.Pow(fn.Sub(x0, x1), 2).Sum().At(0, 0) / float64(x0.Data().Len())
		y := fn.MeanSquaredError(x0, x1)
		assert.True(t, allClose(t, dz.NewVariable(core.New1D(expected)), y))
	})

	t.Run("gradientCheck 1", func(t *testing.T) {

		x0 := dz.NewVariable(core.NewRandN(core.Shape{1, 10}))
		x1 := dz.NewVariable(core.NewRandN(core.Shape{1, 10}))
		f := func(x dz.Variable) dz.Variable {
			return fn.MeanSquaredError(x0, x1)
		}
		assert.True(t, gradientCheck(t, f, x0))
	})

	t.Run("gradientCheck 2", func(t *testing.T) {

		x0 := dz.NewVariable(core.NewRandN(core.Shape{1, 100}))
		x1 := dz.NewVariable(core.NewRandN(core.Shape{1, 100}))
		f := func(x dz.Variable) dz.Variable {
			return fn.MeanSquaredError(x0, x1)
		}
		assert.True(t, gradientCheck(t, f, x0))
	})

	// Skip Test

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		// "success: learn": {
		// 	calc: func() []core.Matrix {
		// 		x := dz.NewVariable(core.NewRand(core.Shape{R: 100, C: 1}), dz.VarOpts().Name("x"))

		// 		a1 := x.Data().CopyApply(func(f float64) float64 { return math.Sin(2 * math.Pi * f) })
		// 		a2 := core.NewRand(core.Shape{R: 100, C: 1})
		// 		y := fn.Add(dz.NewVariable(a1), dz.NewVariable(a2))

		// 		I, H, O := 1, 10, 1
		// 		W1 := fn.MulFloat(dz.NewVariable(core.NewRand(core.Shape{R: I, C: H})), 0.01)
		// 		b1 := dz.NewVariable(core.NewMat(core.Shape{R: 1, C: H}))
		// 		W2 := fn.MulFloat(dz.NewVariable(core.NewRand(core.Shape{R: H, C: O})), 0.01)
		// 		b2 := dz.NewVariable(core.NewMat(core.Shape{R: 1, C: O}))

		// 		y.SetName("y")
		// 		W1.SetName("W1")
		// 		W2.SetName("W2")
		// 		b1.SetName("b1")
		// 		b2.SetName("b2")

		// 		predict := func(x dz.Variable) dz.Variable {
		// 			y := fn.Linear(x, W1, b1)
		// 			y = fn.Sigmoid(y)
		// 			y = fn.Linear(y, W2, b2)
		// 			return y
		// 		}

		// 		lr := 0.2
		// 		iters := 5000
		// 		loss := dz.NewVariable(core.NewEmpty())
		// 		for i := 0; i < iters; i++ {
		// 			ypred := predict(x)
		// 			loss = fn.MeanSquaredError(y, ypred)

		// 			dz.ClearGrad(W1, b1, W2, b2)
		// 			loss.Backward(dz.RetainGrad(true))
		// 			dz.UpdateData(lr, W1, W2, b1, b2)

		// 			if i%100 == 0 {
		// 				t.Log(loss)
		// 			}
		// 		}
		// 		return []core.Matrix{loss.Data()}

		// 	},
		// 	result: []core.Matrix{core.New1D(0.4)},
		// },
	}

	for name, tc := range testCases {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {

			t.Skip("MeanSquarerError Self test")
			c := tc.calc()
			for i := range tc.result {
				assert.Greater(t, tc.result[i].At(0, 0), c[i].At(0, 0))
			}
		})
	}
}
