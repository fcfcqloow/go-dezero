package models

import (
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/ly"
)

type BetterRNN interface {
	Model
	ResetState()
}
type betterRNN struct {
	Model
}

func NewBetterRNN(hiddenSize, outSize int) SimpleRNN {
	instance := new(simpleRNN)
	instance.Model = ExtendsModel(dz.ExtendsLayer(instance.Forward))
	instance.Set("rnn", ly.NewLSTM(hiddenSize, 0))
	instance.Set("fc", ly.NewLinear(outSize))
	return instance
}

func (s *betterRNN) ResetState() {
	s.rnn().ResetState()
}

func (s *betterRNN) Forward(xs ...dz.Variable) dz.Variables {
	x := xs[0]
	h := s.rnn().Apply(x).First()
	y := s.fc().Apply(h)
	return y
}

func (s *betterRNN) rnn() ly.RNN {
	return s.Get("rnn").(ly.RNN)
}

func (s *betterRNN) fc() ly.Linear {
	return s.Get("fc").(ly.Linear)
}
