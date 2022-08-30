package ly

import (
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
)

type LSTM interface {
	dz.Layer
}
type lstm struct {
	dz.Layer
	h, c dz.Variable
}

func NewLSTM(hiddenSize int, inSize int) LSTM {
	H, I := hiddenSize, inSize
	options := []LinearOption{}
	if I != 0 {
		options = append(options, InSize(I))
	}
	instance := new(lstm)
	instance.Layer = dz.ExtendsLayer(instance.Forward)
	instance.Set("x2f", NewLinear(H, options...))
	instance.Set("x2i", NewLinear(H, options...))
	instance.Set("x2o", NewLinear(H, options...))
	instance.Set("x2u", NewLinear(H, options...))

	options = append(options, Nobias(true))
	instance.Set("h2f", NewLinear(H, options...))
	instance.Set("h2i", NewLinear(H, options...))
	instance.Set("h2o", NewLinear(H, options...))
	instance.Set("h2u", NewLinear(H, options...))

	return instance
}
func (l *lstm) ResetState() {
	l.c = nil
	l.h = nil
}
func (l *lstm) Forward(xs ...dz.Variable) dz.Variables {
	x := xs[0]
	var f, i, o, u, cNew, hNew dz.Variable
	fv := l.GetLayer("x2f").Apply(x)[0]
	iv := l.GetLayer("x2i").Apply(x)[0]
	ov := l.GetLayer("x2o").Apply(x)[0]
	uv := l.GetLayer("x2u").Apply(x)[0]
	if l.h == nil {
		f = fn.Sigmoid(fv)
		i = fn.Sigmoid(iv)
		o = fn.Sigmoid(ov)
		u = fn.Tanh(uv)
	} else {
		hfv := l.GetLayer("h2f").Apply(l.h)[0]
		hiv := l.GetLayer("h2i").Apply(l.h)[0]
		hov := l.GetLayer("h2o").Apply(l.h)[0]
		huv := l.GetLayer("h2u").Apply(l.h)[0]
		f = fn.Sigmoid(fn.Add(fv, hfv))
		i = fn.Sigmoid(fn.Add(iv, hiv))
		o = fn.Sigmoid(fn.Add(ov, hov))
		u = fn.Tanh(fn.Add(uv, huv))
	}

	if l.c == nil {
		cNew = fn.Mul(i, u)
	} else {
		cNew = fn.Add(fn.Mul(f, l.c), fn.Mul(i, u))
	}

	hNew = fn.Mul(o, fn.Tanh(cNew))

	l.h, l.c = hNew, cNew
	return []dz.Variable{hNew}
}
