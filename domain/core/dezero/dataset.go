package dz

import "github.com/DolkMd/go-dezero/domain/core"

type (
	datasetOption struct {
		Train           *bool
		Transform       func(core.Matrix) core.Matrix
		TargetTransform func(core.Matrix) core.Matrix
	}
	DatasetOption func(*datasetOption)
)

func ApplyDataSetOpt(options ...DatasetOption) datasetOption {
	option := datasetOption{}
	for _, opt := range options {
		opt(&option)
	}
	return option
}

func TransformData(fn func(core.Matrix) core.Matrix) DatasetOption {
	return func(so *datasetOption) {
		so.Transform = fn
	}
}
func TransformLabel(fn func(core.Matrix) core.Matrix) DatasetOption {
	return func(so *datasetOption) {
		so.TargetTransform = fn
	}
}
func Train(train bool) DatasetOption {
	return func(so *datasetOption) {
		so.Train = &train
	}
}

type (
	Dataset interface {
		Get(idx interface{}) (core.Matrix, core.Matrix)
		Len() int
		IsTrain() bool
	}
	dataset struct {
		data            core.Matrix
		label           core.Matrix
		transform       func(core.Matrix) core.Matrix
		targetTransform func(core.Matrix) core.Matrix
		prepare         func()
		isTrain         bool
	}
)

func ExtendsDataset(data, label core.Matrix, prepare func(), options ...DatasetOption) Dataset {
	option := ApplyDataSetOpt(options...)
	if option.TargetTransform == nil {
		option.TargetTransform = func(x core.Matrix) core.Matrix { return x }
	}
	if option.Transform == nil {
		option.Transform = func(x core.Matrix) core.Matrix { return x }
	}

	instance := &dataset{
		prepare:         prepare,
		data:            data,
		label:           label,
		transform:       option.Transform,
		targetTransform: option.TargetTransform,
	}
	if option.Train != nil {
		instance.isTrain = *option.Train
	}
	if instance.prepare != nil {
		instance.prepare()
	}

	return instance
}

func (d *dataset) Get(idx interface{}) (core.Matrix, core.Matrix) {
	if d.label == nil {
		return d.transform(d.data.Cat(idx)), nil
	}
	return d.transform(d.data.Cat(idx)), d.targetTransform(d.label.Cat(idx))
}

func (d *dataset) Len() int {
	return d.data.Len()
}
func (d *dataset) IsTrain() bool {
	return d.isTrain
}
