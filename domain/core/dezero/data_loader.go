package dz

import (
	"math"

	"github.com/DolkMd/go-dezero/domain/core"
)

type (
	dataLoaderOption struct {
		Shuffle *bool
	}
	DataLoaderOption func(*dataLoaderOption)
	DataLoader       interface {
		Reset()
		Next() bool
		Read() (x, t Variable)
		Len() int
	}
	dataLoader struct {
		dataSet   Dataset
		batchSize int
		shuffle   bool
		dataSize  int
		maxIter   int
		iteration int
		index     []int
		x, t      Variable
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

func NewDataLoader(dataSet Dataset, batchSize int, options ...DataLoaderOption) DataLoader {
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
func (d *dataLoader) Read() (Variable, Variable) {
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
	i, batchSize := d.iteration, d.batchSize
	batchIndex := d.index[i*batchSize : (i+1)*batchSize]
	batchx := []core.Matrix{}
	batcht := []core.Matrix{}
	for _, v := range batchIndex {
		x, t := d.dataSet.Get(v)
		batchx = append(batchx, x)
		batcht = append(batcht, t)
	}
	d.x = NewVariable(core.Merge1DMatrixes(batchx...))
	d.t = NewVariable(core.Merge1DMatrixes(batcht...))

	d.iteration += 1

	return true

}
