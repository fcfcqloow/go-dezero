package fn

import (
	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type (
	softmaxCrossEntropy struct {
		dz.Function
	}
)

func NewSoftmaxCrossEntropy() dz.Function {
	instance := new(softmaxCrossEntropy)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "SoftmaxCrossEntropy")
	return instance
}

func (s *softmaxCrossEntropy) Forward(variables ...dz.Variable) dz.Variables {
	x, t := variables[0], variables[1]
	n := x.Shape().R
	intRancge := []int{}
	for i := 0; i < n; i++ {
		intRancge = append(intRancge, i)
	}
	logz := logsumexp(x)
	logp := Sub(x, logz)
	logp = dz.NewVariable(core.New1D(logp.Data().Search(intRancge, t.Data().Flatten())...))
	a := Neg(dz.NewVariable(logp.Sum()))
	return []dz.Variable{Div(a, dz.NewVariable(core.New1D(float64(n))))}
}

func (s *softmaxCrossEntropy) Backward(variables ...dz.Variable) dz.Variables {
	gy := variables[0]
	x, t := s.Inputs()[0], s.Inputs()[1]
	N, CLS_NUM := x.Shape().R, x.Shape().C

	gy = Mul(gy, dz.NewVariable(core.New1D(1.0/float64(N))))
	y := Softmax(x, core.Axis(1))
	tOnehot := core.Eye(CLS_NUM, CLS_NUM).Cat(t.Data().Flatten())
	a := Sub(y, dz.NewVariable(tOnehot))
	y = Mul(a, gy)
	return []dz.Variable{y, y}
}

func logsumexp(x dz.Variable) dz.Variable {
	m := dz.NewVariable(x.Data().Max(core.Axis(1)))
	y := Sub(x, m)
	y = Exp(y)
	s := dz.NewVariable(y.Sum(core.Axis(1)))
	s = Log(s)
	m = Add(m, s)
	return m
}
