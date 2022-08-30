package dz

import (
	"fmt"
	"sort"

	"github.com/fcfcqloow/go-advance/log"

	"github.com/DolkMd/go-dezero/domain/core"
	"github.com/DolkMd/go-dezero/domain/util"
)

type (
	Variable interface {
		fmt.Stringer
		Data() core.Matrix
		SetData(core.Matrix)
		AddData(core.Matrix)
		SubData(core.Matrix)
		MulData(core.Matrix)
		DivData(core.Matrix)

		Grad() Variable
		SetGrad(Variable)
		ClearGrad()
		Unchain()

		Creator() Function
		SetCreator(Function)

		SetName(string)
		Name() string
		ID() string
		Shape() core.Shape
		NDim() (int, int)
		Size() int
		Generation() int
		Backward(...VariableBackfowardOpt)
		UnchainBackward()

		Reshape(core.Shape)
		T() core.Matrix
		Transpose() core.Matrix
		Sum(...core.Option) core.Matrix
	}
	variable struct {
		core.Matrix
		gard       Variable
		creator    Function
		generation int
		name       string
	}
)

func AsVariable(v interface{}) Variable {
	switch typed := v.(type) {
	case []float64:
		return &variable{Matrix: core.New1D(typed...)}
	case [][]float64:
		return &variable{Matrix: core.New2D(typed)}
	case Variable:
		return typed
	case core.Matrix:
		return NewVariable(typed)
	default:
		panic("unknown type for variable")
	}
}

func NewVariable(data core.Matrix, opts ...VariableOpt) Variable {
	instance := &variable{Matrix: data}
	for _, opt := range opts {
		opt(instance)
	}
	return instance
}

func (v *variable) SetCreator(creator Function) {
	v.generation = creator.Generation() + 1
	v.creator = creator
}
func (v *variable) Backward(opts ...VariableBackfowardOpt) {
	opt := varBackfowardOpts{}
	for _, o := range opts {
		o(&opt)
	}

	if v.gard == nil {
		copied := v.Data().CopyFull(1)
		v.gard = NewVariable(copied)
	}

	creators := []Function{}
	seenSet := util.NewSet()

	addFunc := func(f Function) {
		if !seenSet.Contains(f) {
			creators = append(creators, f)
			seenSet.Add(f)
			sort.Slice(creators, func(i, j int) bool {
				return creators[i].Generation() < creators[j].Generation()
			})
		}
	}
	addFunc(v.creator)

	for len(creators) > 0 {
		creator := creators[len(creators)-1]
		creators = creators[:len(creators)-1]
		gys := creator.Outputs().Grads()
		core.UsingConfig(func() error {
			gxs := creator.Backward(gys...)
			// log.Debug(gputil.Ellipsis(fmt.Sprintf("Bw: %s %p", v.name, v), 30))
			for i := range creator.Inputs() {
				x, gx := creator.Inputs()[i], gxs[i]
				// log.Debug(gputil.Ellipsis(fmt.Sprintf("Bw creator: %s", v.Creator().String()), 50))
				if x.Grad() == nil {
					x.SetGrad(gx)
				} else {
					added := _add(x.Grad(), gx)
					x.SetGrad(added)
				}

				if x.Creator() != nil {
					addFunc(x.Creator())
				}

			}
			return nil
		}, core.EnableBackprop(opt.createGraph))

		if !opt.retainGrad {
			for _, y := range creator.Outputs() {
				log.Debug("clear grad", y.String())
				y.SetGrad(nil)
			}
		}
	}
}
func (v *variable) UnchainBackward() {
	if v.creator != nil {
		creators := []Function{v.creator}
		for len(creators) > 0 {
			creator := creators[len(creators)-1]
			creators = creators[:len(creators)-1]
			for _, input := range creator.Inputs() {
				if input.Creator() != nil {
					creators = append(creators, input.Creator())
					input.Unchain()
				}
			}
		}
	}
}
func (v *variable) ClearGrad()               { v.gard = nil }
func (v *variable) Unchain()                 { v.creator = nil }
func (v *variable) SetData(data core.Matrix) { v.Matrix = data }
func (v *variable) AddData(data core.Matrix) { v.Matrix = v.Matrix.CopyAdd(data) }
func (v *variable) SubData(data core.Matrix) { v.Matrix = v.Matrix.CopySub(data) }
func (v *variable) MulData(data core.Matrix) { v.Matrix = v.Matrix.CopyMul(data) }
func (v *variable) DivData(data core.Matrix) { v.Matrix = v.Matrix.CopyDiv(data) }
func (v *variable) SetName(name string)      { v.name = name }
func (v *variable) SetGrad(dence Variable)   { v.gard = dence }
func (v *variable) Reshape(s core.Shape)     { v.Data().Reshape(s) }
func (v *variable) Data() core.Matrix        { return v.Matrix }

func (v *variable) Grad() Variable                      { return v.gard }
func (v *variable) Name() string                        { return v.name }
func (v *variable) Creator() Function                   { return v.creator }
func (v *variable) Generation() int                     { return v.generation }
func (v *variable) T() core.Matrix                      { return v.Transpose() }
func (v *variable) NDim() (int, int)                    { return v.Data().Dims() }
func (v *variable) Shape() core.Shape                   { return v.Data().Shape() }
func (v *variable) Transpose() core.Matrix              { return v.Data().CopyT() }
func (v *variable) String() string                      { return v.Data().String() + fmt.Sprintf("%p", v) }
func (v *variable) ID() string                          { return fmt.Sprintf("%p", v) }
func (v *variable) Sum(opts ...core.Option) core.Matrix { return v.Data().Sum(opts...) }
func (v *variable) Size() int                           { r, c := v.NDim(); return r * c }

// variable options

type (
	VariableOpt           func(*variable)
	varOpts               struct{}
	varBackfowardOpts     struct{ retainGrad, createGraph bool }
	VariableBackfowardOpt func(*varBackfowardOpts)
)

func VarOpts() varOpts {
	return varOpts{}
}

func (varOpts) Name(name string) VariableOpt {
	return func(v *variable) { v.name = name }
}

func RetainGrad(retainGrad bool) VariableBackfowardOpt {
	return func(opt *varBackfowardOpts) {
		opt.retainGrad = retainGrad
	}
}

func CreateGraph(createGraph bool) VariableBackfowardOpt {
	return func(opt *varBackfowardOpts) {
		opt.createGraph = createGraph
	}
}
