package dz

type (
	OptimizerOption func(*optimizerOption)
	optimizerOption struct {
		Lr           float64
		Momentum     float64
		Alpha        float64
		Beta1, Beta2 float64
		Eps          float64
	}
)

func ApplyOptimizerOption(options ...OptimizerOption) optimizerOption {
	option := optimizerOption{}
	for _, opt := range options {
		opt(&option)
	}

	if option.Lr == 0 {
		option.Lr = 0.01
	}

	if option.Momentum == 0 {
		option.Momentum = 0.9
	}

	if option.Alpha == 0 {
		option.Alpha = 0.001
	}
	if option.Beta1 == 0 {
		option.Beta1 = 0.9
	}
	if option.Beta2 == 0 {
		option.Beta2 = 0.999
	}
	if option.Eps == 0 {
		option.Eps = 1e-8
	}
	return option
}

func Momentum(v float64) OptimizerOption {
	return func(oo *optimizerOption) {
		oo.Momentum = v
	}
}
func Lr(lr float64) OptimizerOption {
	return func(o *optimizerOption) {
		o.Lr = lr
	}
}
func Alpha(alpha float64) OptimizerOption {
	return func(o *optimizerOption) {
		o.Alpha = alpha
	}
}
func Beta1(beta1 float64) OptimizerOption {
	return func(o *optimizerOption) {
		o.Beta1 = beta1
	}
}
func Beta2(beta2 float64) OptimizerOption {
	return func(o *optimizerOption) {
		o.Beta2 = beta2
	}
}

func Eps(eps float64) OptimizerOption {
	return func(o *optimizerOption) {
		o.Eps = eps
	}
}

type (
	Optimizer interface {
		AddHook(f Function)
		UpdateOne(p Variable)
		Update()
		Setup(target Layer) Optimizer
	}
	optimizer struct {
		target    Layer
		hooks     []Function
		updateOne func(Variable)
	}
)

func NewOptimizer(update func(Variable)) Optimizer {
	return &optimizer{
		updateOne: update,
	}
}

func (o *optimizer) Setup(target Layer) Optimizer {
	o.target = target
	return o
}

func (o *optimizer) Update() {
	params := []Variable{}
	for _, p := range o.target.Params() {
		params = append(params, p)
	}

	for _, f := range o.hooks {
		f.Apply(params...)
	}

	for _, p := range params {
		o.UpdateOne(p)
	}

}

func (o *optimizer) UpdateOne(p Variable) {
	o.updateOne(p)
}

func (o *optimizer) AddHook(f Function) {
	o.hooks = append(o.hooks, f)
}
