package optimizers_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/optimizers"
	"github.com/stretchr/testify/assert"
)

func TestSGD(t *testing.T) {
	a := dz.AsVariable([]float64{1, 2, 3})
	a.SetGrad(dz.AsVariable([]float64{1, 2, 3}))
	optimizers.NewSGD(dz.Lr(1)).UpdateOne(a)
	assert.Equal(t, a.Data(), core.New1D(0, 0, 0))
	assert.Equal(t, a.Grad().Data(), core.New1D(1, 2, 3))
}
