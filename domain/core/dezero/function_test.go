package dz_test

import (
	"reflect"
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/stretchr/testify/assert"
)

func TestExtendsFunction(t *testing.T) {
	// fun := dz.ExtendsFunction(nil, nil, "test")
	// assert.IsType(t, (*function)(nil), fun)
}

func Test_function_Apply(t *testing.T) {

	type fields struct {
		forward  dz.Forward
		backward dz.Backward
		name     string
	}
	tests := map[string]struct {
		fields fields
		args   []dz.Variable
		want   dz.Variables
	}{

		"success: x = y": {
			fields: fields{
				forward: func(v ...dz.Variable) dz.Variables {
					return v
				},
				backward: func(v ...dz.Variable) dz.Variables {
					return v
				},
				name: "test",
			},
			args: []dz.Variable{dz.NewVariable(core.New1D(1))},
			want: []dz.Variable{dz.NewVariable(core.New1D(1))},
		},
		// "success: 2x = y": {
		// 	fields: fields{
		// 		forward: func(v ...dz.Variable) dz.Variables {
		// 			return []dz.Variable{fn.MulFloat(v[0], 2)}
		// 		},
		// 		backward: func(v ...dz.Variable) dz.Variables {
		// 			return v
		// 		},
		// 		name: "test",
		// 	},
		// 	args: []dz.Variable{dz.NewVariable(core.New1D(2))},
		// 	want: []dz.Variable{dz.NewVariable(core.New1D(4))},
		// },
		// "success: x**2 = y": {
		// 	fields: fields{
		// 		forward: func(v ...dz.Variable) dz.Variables {
		// 			return []dz.Variable{fn.Pow(v[0], 2)}
		// 		},
		// 		backward: func(v ...dz.Variable) dz.Variables {
		// 			return v
		// 		},
		// 		name: "test",
		// 	},
		// 	args: []dz.Variable{dz.NewVariable(core.New1D(2))},
		// 	want: []dz.Variable{dz.NewVariable(core.New1D(4))},
		// },
	}
	for name, tt := range tests {
		tt := tt
		name := name
		t.Run(name, func(t *testing.T) {

			f := dz.ExtendsFunction(tt.fields.forward, tt.fields.backward, tt.fields.name)
			got := f.Apply(tt.args...)
			equalVariables(t, tt.want, got)
			equalVariables(t, tt.args, f.Inputs())
			equalVariables(t, tt.want, f.Outputs())
			assert.Equal(t, f, f.Outputs().F().Creator())
		})
	}
}

func Test_function_SetInputs(t *testing.T) {
	type fields struct {
		forward  dz.Forward
		backward dz.Backward

		name string
	}
	type args struct {
		vs dz.Variables
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := dz.ExtendsFunction(tt.fields.forward, tt.fields.backward, tt.fields.name)

			f.SetInputs(tt.args.vs)
		})
	}
}

func Test_function_Backward(t *testing.T) {
	type fields struct {
		forward  dz.Forward
		backward dz.Backward

		name string
	}
	type args struct {
		dence []dz.Variable
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   dz.Variables
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := dz.ExtendsFunction(tt.fields.forward, tt.fields.backward, tt.fields.name)

			if got := f.Backward(tt.args.dence...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("function.dz.Backward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_function_Forward(t *testing.T) {
	type fields struct {
		forward  dz.Forward
		backward dz.Backward

		name string
	}
	type args struct {
		dence []dz.Variable
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   dz.Variables
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := dz.ExtendsFunction(tt.fields.forward, tt.fields.backward, tt.fields.name)

			if got := f.Forward(tt.args.dence...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("function.dz.Forward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_function_Generation(t *testing.T) {
	type fields struct {
		forward  dz.Forward
		backward dz.Backward

		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := dz.ExtendsFunction(tt.fields.forward, tt.fields.backward, tt.fields.name)

			if got := f.Generation(); got != tt.want {
				t.Errorf("function.Generation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_function_Name(t *testing.T) {
	type fields struct {
		forward  dz.Forward
		backward dz.Backward

		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := dz.ExtendsFunction(tt.fields.forward, tt.fields.backward, tt.fields.name)

			if got := f.Name(); got != tt.want {
				t.Errorf("function.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_function_String(t *testing.T) {
	type fields struct {
		forward  dz.Forward
		backward dz.Backward

		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := dz.ExtendsFunction(tt.fields.forward, tt.fields.backward, tt.fields.name)

			if got := f.String(); got != tt.want {
				t.Errorf("function.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumericalDiff(t *testing.T) {
	type args struct {
		f   func(dz.Variable) dz.Variable
		x   dz.Variable
		eps *float64
	}
	tests := []struct {
		name string
		args args
		want core.Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dz.NumericalDiff(tt.args.f, tt.args.x, tt.args.eps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumericalDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
