package optimizers_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/optimizers"
	"github.com/stretchr/testify/assert"
)

func TestAdam(t *testing.T) {
	input := [][]float64{{-0.00199303}}
	grad := [][]float64{{-5.19047491}}
	in := dz.AsVariable(core.New2D(input))
	g := dz.AsVariable(core.New2D(grad))
	in.SetGrad(g)
	optimizer := optimizers.NewAdam()
	optimizer.SetT(1)
	optimizer.UpdateOne(in)
	assert.Equal(t, in.Data().At(0, 0), -0.0009930300609246262)
}
