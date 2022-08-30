package optimizers

import (
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
)

type SGD dz.Optimizer
type sgd struct {
	dz.Optimizer
	lr float64
}

func NewSGD(options ...dz.OptimizerOption) SGD {
	option := dz.ApplyOptimizerOption(options...)

	instance := new(sgd)
	instance.Optimizer = dz.NewOptimizer(instance.UpdateOne)
	instance.lr = option.Lr

	return instance
}

func (s *sgd) UpdateOne(param dz.Variable) {
	a1 := fn.MulFloat(param.Grad(), s.lr)
	param.SetData(param.Data().CopySub(a1.Data()))
}
