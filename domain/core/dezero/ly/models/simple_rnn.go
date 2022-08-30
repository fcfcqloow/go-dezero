package models

import (
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/ly"
)

type SimpleRNN interface {
	Model
	ResetState()
}
type simpleRNN struct {
	Model
}

func NewSimpleRNN(hiddenSize, outSize int) SimpleRNN {
	instance := new(simpleRNN)
	instance.Model = ExtendsModel(dz.ExtendsLayer(instance.Forward))
	instance.Set("rnn", ly.NewRNN(hiddenSize))
	instance.Set("fc", ly.NewLinear(outSize))
	return instance
}

func (s *simpleRNN) ResetState() {
	s.rnn().ResetState()
}

func (s *simpleRNN) Forward(xs ...dz.Variable) dz.Variables {
	x := xs[0]
	h := s.rnn().Apply(x).First()
	y := s.fc().Apply(h)
	return y
}

func (s *simpleRNN) rnn() ly.RNN {
	return s.Get("rnn").(ly.RNN)
}

func (s *simpleRNN) fc() ly.Linear {
	return s.Get("fc").(ly.Linear)
}
