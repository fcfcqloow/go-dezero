package models

import (
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/DolkMd/go-dezero/domain/core/dezero/ly"
)

type (
	TwoLayerNet Model
	twoLayerNet struct {
		Model
	}
)

func NewTwoLayerNet(hiddenSize, outSize int, options ...ly.LinearOption) TwoLayerNet {
	instance := new(twoLayerNet)
	instance.Model = ExtendsModel(dz.ExtendsLayer(instance.Forward))
	instance.Set("l1", ly.NewLinear(hiddenSize, options...))
	instance.Set("l2", ly.NewLinear(outSize, options...))

	return instance
}

func (t *twoLayerNet) Forward(xs ...dz.Variable) dz.Variables {
	x := xs[0]
	y := fn.Sigmoid(t.GetLayer("l1").Apply(x).F())
	y = t.GetLayer("l2").Apply(y).First()
	return []dz.Variable{y}
}
