package dz_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/datasets"
	"github.com/DolkMd/go-dezero/domain/core/dezero/loader"
	"github.com/stretchr/testify/assert"
)

func TestDataLoader(t *testing.T) {
	trainSet := datasets.NewSinCurve(dz.Train(true))
	dataloader := loader.NewSeqDataLoader(trainSet, 3)
	dataloader.Next()
	x, _t := dataloader.Read()
	assert.Equal(t, x.Data(), core.New2D([][]float64{
		{-0.02121308873915552},
		{0.9070159825186096},
		{-0.9084922441783305},
	}))
	assert.Equal(t, _t.Data(), core.New2D([][]float64{
		{0.04397207086854972},
		{0.8387896952309626},
		{-0.9155575437357086},
	}))
}
