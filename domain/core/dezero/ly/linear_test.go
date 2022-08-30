package ly_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/ly"
	"github.com/stretchr/testify/assert"
)

func Test_linear_Forward(t *testing.T) {
	type args struct {
		x dz.Variable
	}
	tests := []struct {
		name    string
		outSize int
		args    args
		want    core.Matrix
	}{
		{
			args: args{
				x: dz.NewVariable(core.New1D(1)),
			},
			outSize: 1,
			want:    core.New1D(-0.28158587086436215),
		},
		{
			args: args{
				x: dz.NewVariable(core.New1D(1, 2, 3, 3)),
			},
			outSize: 1,
			want:    core.New1D(-1.2671364188896297),
		},
		{
			args: args{
				x: dz.NewVariable(core.New1D(1, 2, 3, 3)),
			},
			outSize: 3,
			want:    core.New1D(-1.2671364188896297, -1.2671364188896297, -1.2671364188896297),
		},
		{
			args: args{
				x: dz.NewVariable(core.New2D([][]float64{
					{11, 32, 44},
					{1, 3, 5},
					{2, 6, 5},
				})),
			},
			outSize: 10,
			want: core.New2D([][]float64{
				{-14.143910015887517, -14.143910015887517, -14.143910015887517, -14.143910015887517, -14.143910015887517, -14.143910015887517, -14.143910015887517, -14.143910015887517, -14.143910015887517, -14.143910015887517},
				{-1.463163105091812, -1.463163105091812, -1.463163105091812, -1.463163105091812, -1.463163105091812, -1.463163105091812, -1.463163105091812, -1.463163105091812, -1.463163105091812, -1.463163105091812},
				{-2.1134578184659505, -2.1134578184659505, -2.1134578184659505, -2.1134578184659505, -2.1134578184659505, -2.1134578184659505, -2.1134578184659505, -2.1134578184659505, -2.1134578184659505, -2.1134578184659505},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			l := ly.NewLinear(tt.outSize, ly.Seed(0))
			got := l.Forward(tt.args.x)
			if !assert.Equal(t, tt.want, got.F().Data()) {
				t.Log(got)
			}
		})
	}
}
