package ly_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/ly"
	"github.com/stretchr/testify/assert"
)

func TestRNN(t *testing.T) {
	rnn := ly.NewRNN(10)
	x := dz.AsVariable(core.NewRand(core.Shape{R: 1, C: 1}))
	h := rnn.Apply(x).First()
	assert.Equal(t, h.Shape(), core.Shape{R: 1, C: 10})
}
