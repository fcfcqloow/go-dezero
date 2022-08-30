package models_test

import (
	"testing"

	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/ly"
	models "github.com/DolkMd/go-dezero/domain/core/dezero/ly/models"
	"github.com/stretchr/testify/assert"
)

func Test_twoLayerNet_Forward(t *testing.T) {
	tests := []struct {
		name                string
		hiddenSize, outSize int
		want                dz.Variable
		x                   dz.Variable
	}{
		{
			hiddenSize: 2,
			outSize:    3,
			x:          dz.AsVariable([]float64{1, 2, 3}),
			want:       dz.AsVariable([]float64{-0.10903218625537339, -0.10903218625537339, -0.10903218625537339}),
		},
		{
			hiddenSize: 10,
			outSize:    3,
			x: dz.AsVariable([][]float64{
				{1, 2, 3},
				{3, 3, 3},
			}),
			want: dz.AsVariable([][]float64{
				{-0.2438033802024331, -0.2438033802024331, -0.2438033802024331},
				{-0.16739087377636214, -0.16739087377636214, -0.16739087377636214},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			tr := models.NewTwoLayerNet(tt.hiddenSize, tt.outSize, ly.Seed(0))
			got := tr.Forward(tt.x)
			if !assert.Equal(t, tt.want.Data(), got.First().Data()) {
				t.Log(got.First().Data())
			}
		})
	}
}
