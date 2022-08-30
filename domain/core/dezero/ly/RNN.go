package ly

import (
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
)

type (
	RNN interface {
		dz.Layer
		ResetState()
	}
	rnn struct {
		dz.Layer
		h dz.Variable
	}
	rnnOption struct{ inSize *int }
	RNNOption func(*rnnOption)
)

func mergeOption(options ...RNNOption) rnnOption {
	option := rnnOption{}
	for _, opt := range options {
		opt(&option)
	}
	return option
}
func InSizeRnn(inSize int) RNNOption {
	return func(ro *rnnOption) {
		ro.inSize = &inSize
	}
}
func NewRNN(hiddenSize int, options ...RNNOption) RNN {
	option := mergeOption(options...)
	linearOpt := []LinearOption{}
	if option.inSize != nil {
		linearOpt = append(linearOpt, InSize(*option.inSize))
	}

	instance := new(rnn)
	instance.Layer = dz.ExtendsLayer(instance.Forward)
	instance.Set("x2h", NewLinear(hiddenSize, linearOpt...))
	instance.Set("h2h", NewLinear(hiddenSize, append(linearOpt, Nobias(true))...))
	instance.h = nil

	return instance
}

func (r *rnn) Forward(xs ...dz.Variable) dz.Variables {
	x := xs[0]
	var hNew dz.Variable
	if r.h == nil {
		a1 := r.x2h().Apply(x).First()
		hNew = fn.Tanh(a1)
	} else {
		a1 := r.x2h().Apply(x).First()
		a2 := r.h2h().Apply(r.h).First()
		hNew = fn.Tanh(fn.Add(a1, a2))
	}

	r.h = hNew
	return dz.Variables{hNew}
}

func (r *rnn) ResetState() {
	r.h = nil
}

func (r *rnn) x2h() Linear {
	return r.Get("x2h").(*linear)
}

func (r *rnn) h2h() Linear {
	return r.Get("h2h").(*linear)
}
