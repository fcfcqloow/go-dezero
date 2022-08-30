package datasets

import (
	"math/rand"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	appif "github.com/DolkMd/go-dezero/domain/interfaces"
	"github.com/DolkMd/go-dezero/infrastructure/infragraph"
)

type spiral struct{ dz.Dataset }

func NewSpiral(options ...dz.DatasetOption) dz.Dataset {
	opt := dz.ApplyDataSetOpt(options...)
	x, t := GetSpiral(struct{ Train *bool }{opt.Train})
	instance := new(spiral)
	instance.Dataset = dz.ExtendsDataset(x, t, nil, options...)
	return instance
}
func GetSpiralGraphPars(train *bool) []appif.GraphParts {
	x, t := GetSpiral(struct{ Train *bool }{train})
	graph := infragraph.New()
	ps := []appif.GraphParts{}
	for i := 0; i < 3; i++ {
		xs := []float64{}
		ys := []float64{}
		for j := 0; j < x.Shape().R; j++ {
			if int(t.At(0, j)) == i {
				xs = append(xs, x.At(j, 0))
				ys = append(ys, x.At(j, 1))
			}
		}
		colorOpt := appif.GraphColor(255, 255, 0, 255)
		if i == 1 {
			colorOpt = appif.GraphColor(0, 255, 255, 255)
		}
		if i == 2 {
			colorOpt = appif.GraphColor(255, 0, 255, 255)
		}
		p, err := graph.Points(xs, ys, colorOpt)
		if err != nil {
			panic(err)
		}
		ps = append(ps, p)
	}
	return ps
}
func GetSpiral(option struct{ Train *bool }) (x, t core.Matrix) {
	seed := 2020
	if option.Train != nil && *option.Train {
		seed = 1984
	}

	numData, numClass, inputDim := 100.0, 3.0, 2.0
	dataSize := numClass * numData
	x = core.NewMat(core.Shape{R: int(dataSize), C: int(inputDim)})
	t = core.NewMat(core.Shape{R: 1, C: int(dataSize)})

	for i := 0.0; i < numClass; i++ {
		for j := 0.0; j < numData; j++ {
			rate := j / numData
			radius := 1.0 * rate
			theta := fn.AddFloat(fn.MulFloat(
				dz.NewVariable(core.NewRandN(core.Shape{1, 1}, core.RandSeed(float64(seed)))),
				0.2,
			), float64(i*4.0+4.0*rate))
			ix := int(numData*i + j)
			t.Set(0, ix, i)
			x.Set(ix, 0, fn.MulFloat(fn.Sin(theta), radius).Data().At(0, 0))
			x.Set(ix, 1, fn.MulFloat(fn.Cos(theta), radius).Data().At(0, 0))
		}
	}
	x.EachR(func(i int, v []float64) {
		ix := rand.Intn(x.Shape().R)
		x.Swap(i, 0, ix, 0)
		x.Swap(i, 1, ix, 1)
		t.Swap(0, i, 0, ix)
	})

	return
}
