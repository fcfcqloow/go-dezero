package dz_test

import (
	"math"
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/DolkMd/go-dezero/domain/core/dezero/ly"
	"github.com/stretchr/testify/assert"
)

func TestIsLayer(t *testing.T) {
	tests := map[string]struct {
		args interface{}
		want bool
	}{
		"true: layer":        {args: dz.NewLayer(), want: true},
		"false: dz.Variable": {args: dz.NewVariable(nil), want: false},
		"false: int":         {args: 0, want: false},
		"false: string":      {args: "", want: false},
	}
	for name, tt := range tests {
		name := name
		tt := tt
		t.Run(name, func(t *testing.T) {

			assert.Equal(t, tt.want, dz.IsLayer(tt.args))
		})
	}
}

func Test_layer_Apply(t *testing.T) {
	tests := map[string]struct {
		forward func(xs ...dz.Variable) dz.Variables
		args    []dz.Variable
		want    dz.Variables
	}{
		// "success: x*2 (1)": {
		// 	forward: func(xs ...dz.Variable) dz.Variables {
		// 		return []dz.Variable{fn.MulFloat(xs[0], 2)}
		// 	},
		// 	args: []dz.Variable{dz.NewVariable(core.New1D(1))},
		// 	want: []dz.Variable{dz.NewVariable(core.New1D(2))},
		// },
		// "success: x*2 (2)": {
		// 	forward: func(xs ...dz.Variable) dz.Variables {
		// 		return []dz.Variable{fn.MulFloat(xs[0], 2)}
		// 	},
		// 	args: []dz.Variable{dz.NewVariable(core.New1D(3))},
		// 	want: []dz.Variable{dz.NewVariable(core.New1D(6))},
		// },
		// "success: x/2": {
		// 	forward: func(xs ...dz.Variable) dz.Variables {
		// 		return []dz.Variable{fn.DivFloat(xs[0], 2)}
		// 	},
		// 	args: []dz.Variable{dz.NewVariable(core.New1D(10))},
		// 	want: []dz.Variable{dz.NewVariable(core.New1D(5))},
		// },
	}
	for name, tt := range tests {
		tt := tt
		name := name
		t.Run(name, func(t *testing.T) {

			l := dz.ExtendsLayer(tt.forward)
			got := l.Apply(tt.args...)
			equalVariables(t, tt.want, got)
		})
	}
}

func TestLayer(t *testing.T) {

	testCases := map[string]struct {
		calc   func() []core.Matrix
		result []core.Matrix
	}{
		"success: learn": {
			calc: func() []core.Matrix {
				x := dz.NewVariable(core.NewRand(core.Shape{R: 100, C: 1}))
				a1 := x.Data().CopyApply(func(f float64) float64 { return math.Sin(2 * math.Pi * f) })
				a2 := core.NewRand(core.Shape{R: 100, C: 1})
				y := fn.Add(dz.NewVariable(a1), dz.NewVariable(a2))

				l1 := ly.NewLinear(10)
				l2 := ly.NewLinear(1)

				predict := func(x dz.Variable) dz.Variable {
					y := l1.Apply(x).First()
					y = fn.Sigmoid(y)
					y = l2.Apply(y).First()
					return y
				}

				lr := 0.2
				iters := 5000

				loss := dz.NewVariable(core.NewEmpty())
				for i := 0; i < iters; i++ {
					yPred := predict(x)
					loss = fn.MeanSquaredError(y, yPred)

					l1.ClearGrads()
					l2.ClearGrads()
					loss.Backward()

					for _, l := range []dz.Layer{l1, l2} {
						for _, p := range l.Params() {
							a1 := fn.MulFloat(p.Grad(), lr)
							p.SetData(p.Data().CopySub(a1.Data()))
						}
					}
					if i%100 == 0 {
						t.Log(loss)
					}
				}

				return []core.Matrix{loss.Data()}

			},
			result: []core.Matrix{core.New1D(0.4)},
		},
	}

	for name, tc := range testCases {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Skip()

			c := tc.calc()
			for i := range tc.result {
				assert.Greater(t, tc.result[i].At(0, 0), c[i].At(0, 0))
			}
		})
	}
}

// func Test_layer_Params(t *testing.T) {
// 	type fields struct {
// 		params  util.Set
// 		feilds  map[string]interface{}
// 		inputs  dz.Variables
// 		outputs dz.Variables
// 		forward func(xs ...dz.Variable) dz.Variables
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   []Parameter
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &layer{
// 				params:  tt.fields.params,
// 				feilds:  tt.fields.feilds,
// 				inputs:  tt.fields.inputs,
// 				outputs: tt.fields.outputs,
// 				forward: tt.fields.forward,
// 			}
// 			if got := l.Params(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("layer.Params() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_layer_ClearGrads(t *testing.T) {
// 	type fields struct {
// 		params  util.Set
// 		feilds  map[string]interface{}
// 		inputs  dz.Variables
// 		outputs dz.Variables
// 		forward func(xs ...dz.Variable) dz.Variables
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &layer{
// 				params:  tt.fields.params,
// 				feilds:  tt.fields.feilds,
// 				inputs:  tt.fields.inputs,
// 				outputs: tt.fields.outputs,
// 				forward: tt.fields.forward,
// 			}
// 			l.ClearGrads()
// 		})
// 	}
// }

// func Test_layer_Set(t *testing.T) {
// 	type fields struct {
// 		params  util.Set
// 		feilds  map[string]interface{}
// 		inputs  dz.Variables
// 		outputs dz.Variables
// 		forward func(xs ...dz.Variable) dz.Variables
// 	}
// 	type args struct {
// 		name  string
// 		value interface{}
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &layer{
// 				params:  tt.fields.params,
// 				feilds:  tt.fields.feilds,
// 				inputs:  tt.fields.inputs,
// 				outputs: tt.fields.outputs,
// 				forward: tt.fields.forward,
// 			}
// 			l.Set(tt.args.name, tt.args.value)
// 		})
// 	}
// }

// func Test_layer_Get(t *testing.T) {
// 	type fields struct {
// 		params  util.Set
// 		feilds  map[string]interface{}
// 		inputs  dz.Variables
// 		outputs dz.Variables
// 		forward func(xs ...dz.Variable) dz.Variables
// 	}
// 	type args struct {
// 		name string
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   interface{}
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &layer{
// 				params:  tt.fields.params,
// 				feilds:  tt.fields.feilds,
// 				inputs:  tt.fields.inputs,
// 				outputs: tt.fields.outputs,
// 				forward: tt.fields.forward,
// 			}
// 			if got := l.Get(tt.args.name); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("layer.Get() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_layer_GetLayer(t *testing.T) {
// 	type fields struct {
// 		params  util.Set
// 		feilds  map[string]interface{}
// 		inputs  dz.Variables
// 		outputs dz.Variables
// 		forward func(xs ...dz.Variable) dz.Variables
// 	}
// 	type args struct {
// 		name string
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   Layer
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &layer{
// 				params:  tt.fields.params,
// 				feilds:  tt.fields.feilds,
// 				inputs:  tt.fields.inputs,
// 				outputs: tt.fields.outputs,
// 				forward: tt.fields.forward,
// 			}
// 			if got := l.GetLayer(tt.args.name); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("layer.GetLayer() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_layer_GetParameter(t *testing.T) {
// 	type fields struct {
// 		params  util.Set
// 		feilds  map[string]interface{}
// 		inputs  dz.Variables
// 		outputs dz.Variables
// 		forward func(xs ...dz.Variable) dz.Variables
// 	}
// 	type args struct {
// 		name string
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   Parameter
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &layer{
// 				params:  tt.fields.params,
// 				feilds:  tt.fields.feilds,
// 				inputs:  tt.fields.inputs,
// 				outputs: tt.fields.outputs,
// 				forward: tt.fields.forward,
// 			}
// 			if got := l.GetParameter(tt.args.name); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("layer.GetParameter() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_layer_String(t *testing.T) {
// 	type fields struct {
// 		params  util.Set
// 		feilds  map[string]interface{}
// 		inputs  dz.Variables
// 		outputs dz.Variables
// 		forward func(xs ...dz.Variable) dz.Variables
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &layer{
// 				params:  tt.fields.params,
// 				feilds:  tt.fields.feilds,
// 				inputs:  tt.fields.inputs,
// 				outputs: tt.fields.outputs,
// 				forward: tt.fields.forward,
// 			}
// 			if got := l.String(); got != tt.want {
// 				t.Errorf("layer.String() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestFlattenParams(t *testing.T) {
	layer := ly.NewLinear(10)
	dict := map[string]dz.Variable{}
	layer.FlattenParams(dict, "")
	t.Log(dict)

	if err := layer.SaveWeights("test"); err != nil {
		t.Log(err)
	}
	if err := layer.LoadWeights("test"); err != nil {
		t.Log(err)
	}
}
