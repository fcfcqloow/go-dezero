package optimizers

import (
	"fmt"

	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
)

type momentumSgd struct {
	dz.Optimizer
	vs       map[string]dz.Variable
	momentum float64
	lr       float64
}

func NewMomentSGD(options ...dz.OptimizerOption) dz.Optimizer {
	option := dz.ApplyOptimizerOption(options...)

	instance := new(momentumSgd)
	instance.Optimizer = dz.NewOptimizer(instance.UpdateOne)
	instance.lr = option.Lr
	instance.momentum = option.Momentum
	instance.vs = map[string]dz.Variable{}

	return instance
}

func (m *momentumSgd) UpdateOne(param dz.Variable) {
	key := fmt.Sprintf("%p", m)
	if _, ok := m.vs[key]; !ok {
		m.vs[key] = dz.NewVariable(param.Data().CopyFull(0))
	}

	v := m.vs[key]
	v = fn.MulFloat(v, m.momentum)
	a := fn.MulFloat(param.Grad(), m.lr)
	v = fn.Sub(v, a)
	param.SetData(fn.Add(param, v).Data())
}
