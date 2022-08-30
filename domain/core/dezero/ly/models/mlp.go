package models

import (
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/DolkMd/go-dezero/domain/core/dezero/ly"
	cnv "github.com/fcfcqloow/go-advance/convert"
)

type (
	MLP Model
	mlp struct {
		Model
		activation func() Activation
		layers     []dz.Layer
	}
	Activation func(...dz.Variable) dz.Variables
)

func NewMLP(outSizes []int, options ...MlpOption) MLP {
	option := mlpOption{}
	for _, opt := range options {
		opt(&option)
	}

	if option.activation == nil {
		option.activation = func() Activation { return fn.NewSigmoid().Apply }
	}

	instance := new(mlp)
	instance.Model = ExtendsModel(dz.ExtendsLayer(instance.Forward))
	instance.activation = option.activation
	instance.layers = []dz.Layer{}

	lyOptions := []ly.LinearOption{}
	if option.seed != nil {
		lyOptions = append(lyOptions, ly.Seed(*option.seed))
	}
	for i, outSize := range outSizes {
		layer := ly.NewLinear(outSize, lyOptions...)
		instance.Set("l"+cnv.MustStr(i), layer)
		instance.layers = append(instance.layers, layer)
	}
	return instance
}

func (m *mlp) Forward(xs ...dz.Variable) dz.Variables {
	for i := range m.layers[:len(m.layers)-1] {
		tmp := m.layers[i].Apply(xs...)
		xs = m.activation()(tmp...)
	}
	return m.layers[len(m.layers)-1].Apply(xs...)
}

type (
	MlpOption func(*mlpOption)
	mlpOption struct {
		activation func() Activation
		seed       *float64
	}
)

func ActivationFunc(activation Activation) MlpOption {
	return func(m *mlpOption) {
		m.activation = func() Activation {
			return activation
		}
	}
}

func Seed(seed float64) MlpOption {
	return func(m *mlpOption) {
		m.seed = &seed
	}
}
