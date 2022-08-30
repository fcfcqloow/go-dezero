package loader

import (
	"math"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type (
	dataLoaderOption struct {
		Shuffle *bool
	}
	DataLoaderOption func(*dataLoaderOption)
	DataLoader       interface {
		Reset()
		Next() bool
		Read() (x, t dz.Variable)
		Len() int
	}
	dataLoader struct {
		dataSet   dz.Dataset
		batchSize int
		shuffle   bool
		dataSize  int
		maxIter   int
		iteration int
		index     []int
		x, t      dz.Variable
	}
)

func ApplyDataLoaderOption(options ...DataLoaderOption) dataLoaderOption {
	result := dataLoaderOption{}
	for _, opt := range options {
		opt(&result)
	}

	return result
}

func DShuffle(shuffle bool) DataLoaderOption {
	return func(dlo *dataLoaderOption) {
		dlo.Shuffle = &shuffle
	}
}

func NewSeqDataLoader(dataSet dz.Dataset, batchSize int, options ...DataLoaderOption) DataLoader {
	option := ApplyDataLoaderOption(options...)
	instance := new(dataLoader)
	instance.dataSet = dataSet
	instance.dataSize = dataSet.Len()
	instance.batchSize = batchSize
	instance.maxIter = int(math.Ceil(float64(instance.dataSize) / float64(batchSize)))
	if option.Shuffle != nil {
		instance.shuffle = *option.Shuffle
	}
	instance.Reset()

	return instance
}

func (d *dataLoader) Reset() {
	d.iteration = 0
	if d.shuffle {
		d.index = core.RandomPermutation(d.dataSet.Len())
	} else {
		d.index = core.ARange(0, d.dataSet.Len(), 1)
	}
}
func (d *dataLoader) Read() (dz.Variable, dz.Variable) {
	return d.x, d.t
}
func (d *dataLoader) Len() int {
	return d.maxIter
}

func (d *dataLoader) Next() bool {
	if d.iteration >= d.maxIter {
		d.Reset()
		return false
	}
	jump := d.dataSize / d.batchSize
	batchIdx := make([]int, 0, d.batchSize)
	bath := make([][]core.Matrix, 0, d.batchSize)

	for i := 0; i < d.batchSize; i++ {
		batchIdx = append(batchIdx, i*jump+d.iteration)
	}

	for _, i := range batchIdx {
		x, t := d.dataSet.Get(i)
		bath = append(bath, []core.Matrix{x, t})
	}

	batchx := make([]core.Matrix, 0, len(bath))
	batcht := make([]core.Matrix, 0, len(bath))
	for _, b := range bath {
		batchx = append(batchx, b[0])
		batcht = append(batcht, b[1])
	}
	d.x = dz.NewVariable(core.Merge1DMatrixes(batchx...))
	d.t = dz.NewVariable(core.Merge1DMatrixes(batcht...))

	d.iteration += 1

	return true

}
