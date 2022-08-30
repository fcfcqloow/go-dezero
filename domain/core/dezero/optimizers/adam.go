package optimizers

import (
	"math"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
)

type adam struct {
	dz.Optimizer
	t     float64
	beta1 float64
	beta2 float64
	alpha float64
	eps   float64
	ms    map[string]dz.Variable
	vs    map[string]dz.Variable
}

func NewAdam(options ...dz.OptimizerOption) *adam {
	option := dz.ApplyOptimizerOption(options...)
	instance := new(adam)
	instance.Optimizer = dz.NewOptimizer(instance.UpdateOne)
	instance.alpha = option.Alpha
	instance.beta1 = option.Beta1
	instance.beta2 = option.Beta2
	instance.eps = option.Eps
	instance.ms = make(map[string]dz.Variable)
	instance.vs = make(map[string]dz.Variable)
	return instance
}

func (a *adam) Lr() float64 {
	fix1 := 1. - math.Pow(a.beta1, a.t)
	fix2 := 1. - math.Pow(a.beta2, a.t)
	return a.alpha * math.Sqrt(fix2) / fix1
}

func (a *adam) Update() {
	a.t += 1
	a.Optimizer.Update()
}

func (a *adam) Setup(target dz.Layer) dz.Optimizer {
	a.Optimizer = a.Optimizer.Setup(target)
	return a
}

func (a *adam) UpdateOne(param dz.Variable) {
	key := param.ID()
	if _, ok := a.ms[key]; !ok {
		a.ms[key] = dz.AsVariable(core.NewMat(param.Shape()))
		a.vs[key] = dz.AsVariable(core.NewMat(param.Shape()))
	}
	m, v := a.ms[key], a.vs[key]
	beta1, beta2, eps := a.beta1, a.beta2, a.eps
	grad := param.Grad()
	mtmp := fn.MulFloat(fn.Sub(grad, m), 1.-beta1)
	vtmp := fn.MulFloat(fn.Sub(fn.Mul(grad, grad), v), 1.-beta2)
	m.SetData(fn.Add(m, mtmp).Data())
	v.SetData(fn.Add(v, vtmp).Data())
	a1 := fn.MulFloat(m, a.Lr()).Data()
	a2 := v.Data().CopyApply(func(f float64) float64 {
		return math.Sqrt(f) + eps
	})
	param.SetData(param.Data().CopySub(a1.CopyDiv(a2)))
}
