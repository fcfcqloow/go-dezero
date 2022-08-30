package datasets

import (
	"math"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type sinCurve struct {
	dz.Dataset
}

func NewSinCurve(options ...dz.DatasetOption) dz.Dataset {
	opt := dz.ApplyDataSetOpt(options...)
	numData := 1000
	x := core.Linspace(0, 2*math.Pi, numData)
	noiseRange := []float64{-0.05, 0.05}
	noise := core.NewRandUniform(noiseRange[0], noiseRange[1], x.Shape())
	var y core.Matrix
	if opt.Train != nil && *opt.Train == true {
		y = x.CopyApply(math.Sin).CopyAdd(noise)
	} else {
		y = x.CopyApply(math.Cos)
	}
	data := y.Array()[0][:len(y.Array()[0])-1]
	label := y.Array()[0][1:]

	dataMat := core.New1D(data...).CopyT()
	labelMat := core.New1D(label...).CopyT()
	instance := new(sinCurve)
	instance.Dataset = dz.ExtendsDataset(dataMat, labelMat, nil, options...)
	return instance
}
