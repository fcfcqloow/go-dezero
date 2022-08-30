package ly

import (
	"fmt"
	"math"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	cnv "github.com/fcfcqloow/go-advance/convert"
)

const (
	w = "w"
	b = "b"
)

type (
	Linear dz.Layer
	linear struct {
		dz.Layer
		inSize, outSize int
		seed            core.RandSeedOption
	}
	lineOpt struct {
		nobias     *bool
		inSize     *int
		customSeed *float64
	}
	LinearOption func(*lineOpt)
)

func Nobias(n bool) LinearOption {
	return func(lo *lineOpt) {
		lo.nobias = &n
	}
}
func InSize(i int) LinearOption {
	return func(lo *lineOpt) {
		lo.inSize = &i
	}
}
func Seed(seed float64) LinearOption {
	return func(lo *lineOpt) {
		lo.customSeed = &seed
	}
}

func NewLinear(outSize int, opts ...LinearOption) Linear {
	option := lineOpt{}
	for _, opt := range opts {
		opt(&option)
	}

	instance := &linear{outSize: outSize}
	instance.Layer = dz.ExtendsLayer(instance.Forward)
	instance.setW(dz.NewParameter(core.NewEmpty(), dz.VarOpts().Name(cnv.MustStr(outSize))))

	if option.customSeed != nil {
		instance.seed = core.RandSeed(*option.customSeed)
	}

	if option.inSize != nil {
		instance.inSize = *option.inSize
		instance.initW()
	}

	if option.nobias == nil || !*option.nobias {
		instance.setB(dz.NewParameter(core.NewFull(core.Shape{R: 1, C: outSize}, 0), dz.VarOpts().Name("b: out->"+cnv.MustStr(outSize))))
	}

	return instance
}

func (l *linear) initW() {
	I, O := l.inSize, l.outSize
	var Wdata core.Matrix
	if l.seed != nil {
		Wdata = core.NewRandN(core.Shape{R: I, C: O}, l.seed)
	} else {
		Wdata = core.NewRandN(core.Shape{R: I, C: O})
	}

	l.w().SetData(Wdata.CopyApply(func(f float64) float64 {
		return math.Sqrt(1./float64(I)) * f
	}))
}

func (l *linear) Forward(xs ...dz.Variable) dz.Variables {
	x := xs[0]
	if l.w().Shape() == (core.Shape{R: 0, C: 0}) {
		l.inSize = x.Shape().C
		l.initW()
	}
	return []dz.Variable{fn.Linear(x, l.w(), l.b())}
}

func (l *linear) setW(p dz.Parameter) {
	p.SetName(fmt.Sprintf("w: %p", p))
	l.Layer.Set(w, p)
}

func (l *linear) setB(p dz.Parameter) {
	p.SetName(fmt.Sprintf("b: %p", p))
	l.Layer.Set(b, p)
}

func (l *linear) w() dz.Parameter {
	return l.Get(w).(dz.Parameter)
}

func (l *linear) b() dz.Parameter {
	result := l.Get(b)
	if result != nil {
		return result.(dz.Parameter)
	}

	return nil
}
