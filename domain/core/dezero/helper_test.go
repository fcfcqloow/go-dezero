package dz_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/stretchr/testify/assert"
)

func random(t *testing.T, min, max float64) float64 {
	t.Helper()
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(max-min) + min
}

func equalVariables(t *testing.T, x, y dz.Variables) {
	t.Helper()
	for i := range x {
		assert.Equal(t, x[i].Data(), y[i].Data())
	}
}
func allClose(t *testing.T, x, y dz.Variable) bool {
	t.Helper()
	return core.Allclose(x.Data(), y.Data(), nil, nil)
}

func gradientCheck(t *testing.T, f func(dz.Variable) dz.Variable, x dz.Variable) bool {
	t.Helper()

	numGrad := core.NumericalGradient(func(m core.Matrix) core.Matrix {
		return f(dz.NewVariable(m)).Data()
	}, x.Data())
	y := f(x)
	y.Backward()
	bpGrad := x.Grad().Data()

	if numGrad.Shape() != bpGrad.Shape() {
		t.Fatal("gradientCheck: missing shape")
	}
	return allClose(t, dz.NewVariable(numGrad), dz.NewVariable(bpGrad))
}
