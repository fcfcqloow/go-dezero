package fn_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	. "github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestSphere(t *testing.T) {

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		"success: backward": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New1D(1))
				y := dz.NewVariable(core.New1D(1))
				z := Sphere(x, y)
				z.Backward()
				return []core.Matrix{x.Grad().Data(), y.Grad().Data()}
			},
			result: []core.Matrix{core.New1D(2), core.New1D(2)},
		},
		"success: apply": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New1D(3))
				y := dz.NewVariable(core.New1D(3))
				z := Sphere(x, y)
				return []core.Matrix{z.Data()}
			},
			result: []core.Matrix{core.New1D(18)},
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

func TestMatyas(t *testing.T) {

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		"success: backward": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New1D(1))
				y := dz.NewVariable(core.New1D(1))
				z := Matyas(x, y)
				z.Backward()
				return []core.Matrix{x.Grad().Data(), y.Grad().Data()}
			},
			result: []core.Matrix{core.New1D(0.040000000000000036), core.New1D(0.040000000000000036)},
		},
		"success: apply": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New1D(1))
				y := dz.NewVariable(core.New1D(1))
				z := Matyas(x, y)
				return []core.Matrix{z.Data()}
			},
			result: []core.Matrix{core.New1D(0.040000000000000036)},
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

func TestGoldstein(t *testing.T) {

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		"success: backward": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.New1D(1))
				y := dz.NewVariable(core.New1D(1))
				z := Goldstein(x, y)
				z.Backward()
				return []core.Matrix{x.Grad().Data(), y.Grad().Data()}
			},
			result: []core.Matrix{core.New1D(-5376), core.New1D(8064)},
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

func TestRosenbrock(t *testing.T) {

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		"success: backward": {
			calc: func() []core.Matrix {
				x0 := dz.NewVariable(core.New1D(0))
				x1 := dz.NewVariable(core.New1D(2))
				y := Rosenbrock(x0, x1)
				y.Backward()
				return []core.Matrix{x0.Grad().Data(), x1.Grad().Data()}
			},
			result: []core.Matrix{core.New1D(-2), core.New1D(400)},
		},
		"success: learning": {
			calc: func() []core.Matrix {
				x0 := dz.NewVariable(core.New1D(0.0))
				x1 := dz.NewVariable(core.New1D(2.0))
				lr := 0.001 // 学習率
				iter := 10000
				for i := 0; i < iter; i++ {
					y := Rosenbrock(x0, x1)
					x0.ClearGrad()
					x1.ClearGrad()
					y.Backward()

					a := MulFloat(x0.Grad(), lr)
					b := MulFloat(x1.Grad(), lr)

					x0.SetData(Sub(x0, a).Data())
					x1.SetData(Sub(x1, b).Data())
				}
				return []core.Matrix{x0.Data(), x1.Data()}
			},
			result: []core.Matrix{core.New1D(0.9944984367782456), core.New1D(0.9890050527419593)},
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
