package fn_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestRelu(t *testing.T) {
	t.Run("allclose 1", func(t *testing.T) {

		x := core.New2D([][]float64{
			{-1, 0},
			{2, -3},
			{-2, 1},
		})
		ans := core.New2D([][]float64{{0, 0}, {2, 0}, {0, 1}})
		assert.Equal(t, fn.Relu(dz.AsVariable(x)).Data(), ans)
	})
}
