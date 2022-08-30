package core_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	"github.com/stretchr/testify/assert"
)

func TestShape_String(t *testing.T) {

	tests := map[string]struct {
		R    int
		C    int
		want string
	}{
		"Row 10, Col 100": {
			R:    10,
			C:    100,
			want: "(rows: 10, cols: 100)",
		},
		"Row 1, Col 100": {
			R:    1,
			C:    100,
			want: "(rows: 1, cols: 100)",
		},
		"Row 0, Col 0": {
			R:    0,
			C:    0,
			want: "(rows: 0, cols: 0)",
		},
	}
	for name, tt := range tests {
		tt := tt
		name := name
		t.Run(name, func(t *testing.T) {

			s := core.Shape{R: tt.R, C: tt.C}
			if got := s.String(); got != tt.want {
				t.Errorf("Shape.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithShape(t *testing.T) {

	type args struct {
		r int
		c int
	}
	tests := []struct {
		name string
		args args
		want core.Matrix
	}{
		{
			name: "2x2 with shape",
			args: args{r: 2, c: 2},
			want: core.New2D([][]float64{{0, 0}, {0, 0}}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			if got := core.WithShape(tt.args.r, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithShape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMat(t *testing.T) {

	type args struct {
		s core.Shape
	}
	tests := []struct {
		name string
		args args
		want core.Matrix
	}{
		{
			args: args{s: core.Shape{2, 2}},
			want: core.New2D([][]float64{{0, 0}, {0, 0}}),
		},
		{
			args: args{s: core.Shape{1, 1}},
			want: core.New2D([][]float64{{0}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := core.NewMat(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFull(t *testing.T) {

	type args struct {
		shape core.Shape
		v     float64
	}
	tests := []struct {
		name string
		args args
		want core.Matrix
	}{
		{
			name: "2x2 full 2",
			args: args{
				shape: core.Shape{R: 2, C: 2},
				v:     2,
			},
			want: core.New2D([][]float64{
				{2, 2},
				{2, 2},
			}),
		},
		{
			name: "2x3 full 2",
			args: args{
				shape: core.Shape{R: 2, C: 3},
				v:     3,
			},
			want: core.New2D([][]float64{
				{3, 3, 3},
				{3, 3, 3},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := core.NewFull(tt.args.shape, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew1D(t *testing.T) {

	type args struct {
		dence []float64
	}
	tests := []struct {
		name        string
		args        args
		wantFlatten []float64
	}{
		{
			name: "0 array",
			args: args{
				dence: []float64{0, 0, 0, 0},
			},
			wantFlatten: []float64{0, 0, 0, 0},
		},
		{
			name: "any array",
			args: args{
				dence: []float64{2, 20, 10, 91},
			},
			wantFlatten: []float64{2, 20, 10, 91},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			got := core.New1D(tt.args.dence...)
			assert.Equal(t, tt.wantFlatten, got.Flatten())
		})
	}
}

func TestNew2D(t *testing.T) {

	type args struct {
		dence [][]float64
	}
	tests := []struct {
		name        string
		args        args
		wantFlatten []float64
	}{
		{
			args: args{
				dence: [][]float64{
					{1, 2, 3},
				},
			},
			wantFlatten: []float64{1, 2, 3},
		},
		{
			args: args{
				dence: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			wantFlatten: []float64{1, 2, 3, 4, 5, 6},
		},
		{
			args: args{
				dence: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			wantFlatten: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			got := core.New2D(tt.args.dence)
			assert.Equal(t, tt.wantFlatten, got.Flatten())
		})
	}
}

func TestEye(t *testing.T) {

	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		args args
		want core.Matrix
	}{
		{
			args: args{row: 3, col: 3},
			want: core.New2D([][]float64{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			}),
		},
		{
			args: args{row: 4, col: 4},
			want: core.New2D([][]float64{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
			}),
		},
		{
			args: args{row: 4, col: 2},
			want: core.New2D([][]float64{
				{1, 0},
				{0, 1},
				{0, 0},
				{0, 0},
			}),
		},
		{
			args: args{row: 2, col: 4},
			want: core.New2D([][]float64{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			if got := core.Eye(tt.args.row, tt.args.col); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Eye() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matrix_Set(t *testing.T) {

	type args struct {
		mat core.Matrix
		i   int
		j   int
		v   float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				i: 0, j: 1, v: 10,
				mat: core.New1D(10, 99),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			tt.args.mat.Set(tt.args.i, tt.args.j, tt.args.v)
			assert.Equal(t, tt.args.v, tt.args.mat.At(tt.args.i, tt.args.j))
		})
	}
}

// func Test_matrix_Reshape(t *testing.T) {
// 	type args struct {
// 	}
// 	tests := []struct {
// 		name  string
// 		mat   core.Matrix
// 		want  core.Matrix
// 		shape core.Shape
// 	}{
// 		{
// 			mat: core.New2D([][]float64{
// 				{1},
// 				{2},
// 				{3},
// 				{4},
// 				{4},
// 				{5},
// 				{6},
// 			}),
// 			want:  core.New2D([][]float64{}),
// 			shape: core.Shape{2, 2},
// 		},
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
//
// 			tt.mat.Reshape(tt.shape)
// 			assert.Equal(t, tt.mat, tt.want)
// 		})
// 	}
// }

// func Test_matrix_Shape(t *testing.T) {
// 	type fields struct {
// 		Dense *mat.Dense
// 	}
// 	tests := []struct {
// 		name       string
// 		fields     fields
// 		wantResult Shape
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			m := &matrix{
// 				Dense: tt.fields.Dense,
// 			}
// 			if gotResult := m.Shape(); !reflect.DeepEqual(gotResult, tt.wantResult) {
// 				t.Errorf("matrix.Shape() = %v, want %v", gotResult, tt.wantResult)
// 			}
// 		})
// 	}
// }

func Test_matrix_Flatten(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		want []float64
	}{
		{
			mat:  core.New1D(1, 2, 3),
			want: []float64{1, 2, 3},
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{9, 9, 9},
			}),
			want: []float64{1, 2, 3, 9, 9, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mat.Flatten()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_matrix_Dims(t *testing.T) {

	tests := []struct {
		name    string
		mat     core.Matrix
		wantRow int
		wantCol int
	}{
		{
			mat:     core.WithShape(2, 2),
			wantRow: 2,
			wantCol: 2,
		},
		{
			mat:     core.WithShape(10, 2),
			wantRow: 10,
			wantCol: 2,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			gotRow, gotCol := m.Dims()
			if gotRow != tt.wantRow {
				t.Errorf("matrix.Dims() gotRow = %v, want %v", gotRow, tt.wantRow)
			}
			if gotCol != tt.wantCol {
				t.Errorf("matrix.Dims() gotCol = %v, want %v", gotCol, tt.wantCol)
			}
		})
	}
}

func Test_matrix_Rows(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		want int
	}{
		{
			mat:  core.WithShape(10, 1),
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mat
			if got := m.Rows(); got != tt.want {
				t.Errorf("matrix.Rows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matrix_Cols(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		want int
	}{
		{
			mat:  core.WithShape(10, 100),
			want: 100,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			if got := m.Cols(); got != tt.want {
				t.Errorf("matrix.Cols() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matrix_Copy(t *testing.T) {

	mat := core.NewEmpty()
	assert.NotEqual(t, fmt.Sprintf("%p", mat), fmt.Sprintf("%p", mat.Copy()))
}

func Test_matrix_CopyT(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		want core.Matrix
	}{
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			}),
			want: core.New2D([][]float64{
				{1, 1, 1},
				{2, 2, 2},
				{3, 3, 3},
			}),
		},
		{
			mat: core.New2D([][]float64{
				{1},
				{2},
				{3},
			}),
			want: core.New1D(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mat
			got := m.CopyT()
			assert.Equal(t, tt.want, got)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))

		})
	}
}

func Test_matrix_CopyAdd(t *testing.T) {

	tests := []struct {
		name     string
		mat, arg core.Matrix
		want     core.Matrix
	}{
		{
			mat:  core.New1D(1, 2, 3),
			arg:  core.New1D(5, 5, 5),
			want: core.New1D(6, 7, 8),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: core.NewFull(core.Shape{3, 3}, 3),
			want: core.New2D([][]float64{
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.CopyAdd(tt.arg)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			assert.Equal(t, tt.want, got)

		})
	}
}

func Test_matrix_CopySub(t *testing.T) {

	tests := []struct {
		name     string
		mat, arg core.Matrix
		want     core.Matrix
	}{
		{
			mat:  core.New1D(1, 2, 3),
			arg:  core.New1D(5, 5, 5),
			want: core.New1D(-4, -3, -2),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: core.NewFull(core.Shape{3, 3}, 3),
			want: core.New2D([][]float64{
				{-2, -1, 0},
				{1, 2, 3},
				{4, 5, 6},
			}),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: core.New2D([][]float64{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			}),
			want: core.New2D([][]float64{
				{0, 0, 0},
				{3, 3, 3},
				{6, 6, 6},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.CopySub(tt.arg)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			assert.Equal(t, tt.want, got)

		})
	}
}

func Test_matrix_CopyMul(t *testing.T) {

	tests := []struct {
		name     string
		mat, arg core.Matrix
		want     core.Matrix
	}{
		{
			mat:  core.New1D(1, 2, 3),
			arg:  core.New1D(5, 5, 5),
			want: core.New1D(5, 10, 15),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: core.NewFull(core.Shape{3, 3}, 3),
			want: core.New2D([][]float64{
				{3, 6, 9},
				{12, 15, 18},
				{21, 24, 27},
			}),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: core.New2D([][]float64{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			}),
			want: core.New2D([][]float64{
				{1, 4, 9},
				{4, 10, 18},
				{7, 16, 27},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.CopyMul(tt.arg)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_matrix_CopyDiv(t *testing.T) {

	tests := []struct {
		name     string
		mat, arg core.Matrix
		want     core.Matrix
	}{
		{
			mat:  core.New1D(1, 2, 3),
			arg:  core.New1D(5, 5, 5),
			want: core.New1D(0.2, 0.4, 3.0/5.0),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: core.NewFull(core.Shape{3, 3}, 3),
			want: core.New2D([][]float64{
				{1. / 3., 2. / 3., 3. / 3.},
				{4. / 3., 5. / 3., 6. / 3.},
				{7. / 3., 8. / 3., 9. / 3.},
			}),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: core.New2D([][]float64{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			}),
			want: core.New2D([][]float64{
				{1, 1, 1},
				{4, 2.5, 2},
				{7, 4, 3},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.CopyDiv(tt.arg)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_matrix_CopyAddFloat(t *testing.T) {
	tests := []struct {
		name string
		mat  core.Matrix
		arg  float64
		want core.Matrix
	}{
		{
			mat:  core.New1D(1, 2, 3),
			arg:  10,
			want: core.New1D(11, 12, 13),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: 10,
			want: core.New2D([][]float64{
				{11, 12, 13},
				{14, 15, 16},
				{17, 18, 19},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.CopyAddFloat(tt.arg)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			assert.Equal(t, tt.want, got)

		})
	}
}

func Test_matrix_CopySubFloat(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		arg  float64
		want core.Matrix
	}{
		{
			mat:  core.New1D(1, 2, 3),
			arg:  10,
			want: core.New1D(-9, -8, -7),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: 11,
			want: core.New2D([][]float64{
				{-10, -9, -8},
				{-7, -6, -5},
				{-4, -3, -2},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.CopySubFloat(tt.arg)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_matrix_CopyMulFloat(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		arg  float64
		want core.Matrix
	}{
		{
			mat:  core.New1D(1, 2, 3),
			arg:  10,
			want: core.New1D(10, 20, 30),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: 11,
			want: core.New2D([][]float64{
				{11, 22, 33},
				{44, 55, 66},
				{77, 88, 99},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.CopyMulFloat(tt.arg)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_matrix_CopyDivFloat(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		arg  float64
		want core.Matrix
	}{
		{
			mat:  core.New1D(1, 2, 3),
			arg:  10,
			want: core.New1D(0.1, 0.2, 0.3),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: 2,
			want: core.New2D([][]float64{
				{.5, 1, 1.5},
				{2, 2.5, 3},
				{3.5, 4, 4.5},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.CopyDivFloat(tt.arg)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_matrix_CopyMatMul(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		arg  core.Matrix
		want core.Matrix
	}{
		{
			mat: core.New1D(1, 2, 3),
			arg: core.New2D([][]float64{
				{1},
				{2},
				{3},
			}),
			want: core.New1D(14),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: core.New2D([][]float64{
				{.5, 1, 1.5},
				{2, 2.5, 3},
				{3.5, 4, 4.5},
			}),
			want: core.New2D([][]float64{
				{15, 18, 21},
				{33, 40.5, 48},
				{51, 63, 75},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.CopyMatMul(tt.arg)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			if !assert.Equal(t, tt.want, got) {
				t.Log(got)
			}
		})
	}
}

// func Test_matrix_CopyScale(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		mat  core.Matrix
// 		arg  float64
// 		want core.Matrix
// 	}{
// 		{
// 			mat:  core.New1D(1, 2, 3),
// 			arg:  1,
// 			want: core.New1D(14),
// 		},
// 		{
// 			mat: core.New2D([][]float64{
// 				{1, 2, 3},
// 				{4, 5, 6},
// 				{7, 8, 9},
// 			}),
// 			arg: 10,
// 			want: core.New2D([][]float64{
// 				{15, 18, 21},
// 				{33, 40.5, 48},
// 				{51, 63, 75},
// 			}),
// 		},
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
//
// 			m := tt.mat
// 			got := m.CopyScale(tt.arg)
// 			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
// 			if !assert.Equal(t, tt.want, got) {
// 				t.Log(got)
// 			}
// 		})
// 	}
// }

func Test_matrix_CopyLog(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		want core.Matrix
	}{
		{
			mat:  core.New1D(1, 2, 3),
			want: core.New1D(0, 0.6931471805599453, 1.0986122886681096),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			want: core.New2D([][]float64{
				{0, 0.6931471805599453, 1.0986122886681096},
				{1.3862943611198906, 1.6094379124341003, 1.791759469228055},
				{1.9459101490553132, 2.0794415416798357, 2.1972245773362196},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.CopyLog()
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			if !assert.Equal(t, tt.want, got) {
				t.Log(got)
			}
		})
	}
}

func Test_matrix_CopyClip(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		min  float64
		max  float64
		want core.Matrix
	}{
		{
			mat:  core.New1D(1, 2, 3),
			max:  1,
			min:  0,
			want: core.New1D(1, 1, 1),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			max: 7,
			min: 2,
			want: core.New2D([][]float64{
				{2, 2, 3},
				{4, 5, 6},
				{7, 7, 7},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.CopyClip(tt.min, tt.max)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			if !assert.Equal(t, tt.want, got) {
				t.Log(got)
			}
		})
	}
}

func Test_matrix_CopyFull(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		arg  float64
		want core.Matrix
	}{
		{
			mat:  core.New1D(1, 2, 3),
			arg:  2,
			want: core.New1D(2, 2, 2),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: 99,
			want: core.New2D([][]float64{
				{99, 99, 99},
				{99, 99, 99},
				{99, 99, 99},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.CopyFull(tt.arg)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			if !assert.Equal(t, tt.want, got) {
				t.Log(got)
			}
		})
	}
}

func Test_matrix_CopyApply(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		arg  func(float64) float64
		want core.Matrix
	}{
		{
			mat: core.New1D(1, 2, 3),
			arg: func(f float64) float64 {
				return f * 2
			},
			want: core.New1D(2, 4, 6),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: func(f float64) float64 {
				return 2
			},
			want: core.New2D([][]float64{
				{2, 2, 2},
				{2, 2, 2},
				{2, 2, 2},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.CopyApply(tt.arg)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			if !assert.Equal(t, tt.want, got) {
				t.Log(got)
			}
		})
	}
}

// func Test_matrix_CopyReshape(t *testing.T) {
// 	type fields struct {
// 		Dense *mat.Dense
// 	}
// 	type args struct {
// 		s Shape
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   core.Matrix
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			m := &matrix{
// 				Dense: tt.fields.Dense,
// 			}
// 			if got := m.CopyReshape(tt.args.s); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("matrix.CopyReshape() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_matrix_Apply(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		arg  func(float64) float64
		want core.Matrix
	}{
		{
			mat: core.New1D(1, 2, 3),
			arg: func(f float64) float64 {
				return f * 2
			},
			want: core.New1D(2, 4, 6),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: func(f float64) float64 {
				return 2
			},
			want: core.New2D([][]float64{
				{2, 2, 2},
				{2, 2, 2},
				{2, 2, 2},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			got := tt.mat
			got.Apply(tt.arg)
			if !assert.Equal(t, tt.want, got) {
				t.Log(got)
			}
		})
	}
}

func Test_matrix_Len(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		want int
	}{
		{
			mat:  core.New1D(1, 2, 3),
			want: 3,
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{7, 8, 9},
			}),
			want: 4,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.Len()
			if !assert.Equal(t, tt.want, got) {
				t.Log(got)
			}
		})
	}
}

func Test_matrix_Cat(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		arg  interface{}
		want core.Matrix
	}{
		{
			mat:  core.New1D(1, 2, 3),
			arg:  []int{0, 1},
			want: core.New1D(1, 2),
		},
		{
			mat:  core.New1D(1, 2, 3),
			arg:  2.0,
			want: core.New1D(3),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: []int{0, 2},
			want: core.New2D([][]float64{
				{1, 2, 3},
				{7, 8, 9},
			}),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: 2,
			want: core.New2D([][]float64{
				{7, 8, 9},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.Cat(tt.arg)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			if !assert.Equal(t, tt.want, got) {
				t.Log(got)
			}
		})
	}
}

func Test_matrix_Sum(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		arg  []core.Option
		want core.Matrix
	}{
		{
			mat:  core.New1D(1, 2, 3),
			want: core.New1D(6),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: []core.Option{core.Axis(0)},
			want: core.New2D([][]float64{
				{12, 15, 18},
			}),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: []core.Option{core.Axis(1)},
			want: core.New2D([][]float64{
				{6},
				{15},
				{24},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.Sum(tt.arg...)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			if !assert.Equal(t, tt.want, got) {
				t.Log(got)
			}
		})
	}
}

func Test_matrix_Max(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		arg  []core.Option
		want core.Matrix
	}{
		{
			mat:  core.New1D(1, 2, 3),
			want: core.New1D(3),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: []core.Option{core.Axis(0)},
			want: core.New2D([][]float64{
				{7, 8, 9},
			}),
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: []core.Option{core.Axis(1)},
			want: core.New2D([][]float64{
				{3},
				{6},
				{9},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.Max(tt.arg...)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			if !assert.Equal(t, tt.want, got) {
				t.Log(got)
			}
		})
	}
}

func Test_matrix_Allclose(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		arg  core.Matrix
		want bool
	}{
		{
			mat:  core.New1D(1, 2, 3),
			arg:  core.New1D(-1, 2, 3),
			want: false,
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: core.New2D([][]float64{
				{1.000001, 2.000001, 3.000001},
				{4.000001, 5.000001, 6.000001},
				{7.000001, 8.000001, 9.000001},
			}),
			want: true,
		},
		{
			mat: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}),
			arg: core.New2D([][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{70, 8, 9},
			}),
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.Allclose(tt.arg, nil, nil)
			if !assert.Equal(t, tt.want, got) {
				t.Log(got)
			}
		})
	}
}

func Test_matrix_Broadcast(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		arg  core.Shape
		want core.Matrix
	}{
		{
			mat: core.New1D(1, 2, 3),
			arg: core.Shape{R: 3, C: 3},
			want: core.New2D([][]float64{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			}),
		},
		{
			mat: core.New1D(9),
			arg: core.Shape{R: 3, C: 3},
			want: core.New2D([][]float64{
				{9, 9, 9},
				{9, 9, 9},
				{9, 9, 9},
			}),
		},
		{
			mat: core.New2D([][]float64{
				{1},
				{2},
				{3},
			}),
			arg: core.Shape{R: 3, C: 3},
			want: core.New2D([][]float64{
				{1, 1, 1},
				{2, 2, 2},
				{3, 3, 3},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.Broadcast(tt.arg)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			if !assert.Equal(t, tt.want, got) {
				t.Log(got)
			}
		})
	}
}

// func Test_matrix_Each(t *testing.T) {
// 	type fields struct {
// 		Dense *mat.Dense
// 	}
// 	type args struct {
// 		fn func(i, j int, v float64)
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
// 			m := &matrix{
// 				Dense: tt.fields.Dense,
// 			}
// 			m.Each(tt.args.fn)
// 		})
// 	}
// }

// func Test_matrix_EachR(t *testing.T) {
// 	type fields struct {
// 		Dense *mat.Dense
// 	}
// 	type args struct {
// 		fn func(i int, v []float64)
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
// 			m := &matrix{
// 				Dense: tt.fields.Dense,
// 			}
// 			m.EachR(tt.args.fn)
// 		})
// 	}
// }

// func Test_matrix_EachC(t *testing.T) {
// 	type fields struct {
// 		Dense *mat.Dense
// 	}
// 	type args struct {
// 		fn func(i int, v []float64)
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
// 			m := &matrix{
// 				Dense: tt.fields.Dense,
// 			}
// 			m.EachC(tt.args.fn)
// 		})
// 	}
// }

func Test_matrix_OnOff(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		arg  func(i, j int, v float64) bool
		want core.Matrix
	}{
		{
			mat: core.New1D(1, 2, 3),
			arg: func(i, j int, v float64) bool {
				return false
			},
			want: core.New2D([][]float64{
				{0, 0, 0},
			}),
		},
		{
			mat: core.New2D([][]float64{
				{1, 5, 6},
				{2, 9, 2},
				{3, 3, 3},
			}),
			arg: func(i, j int, v float64) bool {
				return v > 3
			},
			want: core.New2D([][]float64{
				{0, 1, 1},
				{0, 1, 0},
				{0, 0, 0},
			}),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.OnOff(tt.arg)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			if !assert.Equal(t, tt.want, got) {
				t.Log(got)
			}
		})
	}
}

// func Test_matrix_Clip(t *testing.T) {
//
// 	tests := []struct {
// 		name     string
// 		mat      core.Matrix
// 		min, max float64
// 		want     core.Matrix
// 	}{
// 		{
// 			mat: core.New1D(1, 2, 3),
// 			want: core.New2D([][]float64{
// 				{0, 0, 0},
// 			}),
// 		},
// 		{
// 			mat: core.New2D([][]float64{
// 				{1, 5, 6},
// 				{2, 9, 2},
// 				{3, 3, 3},
// 			}),
// 			want: core.New2D([][]float64{
// 				{0, 1, 1},
// 				{0, 1, 0},
// 				{0, 0, 0},
// 			}),
// 		},
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
//
// 			m := tt.mat
// 			got := m.Clip(tt.min, tt.max)
// 			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
// 			if !assert.Equal(t, tt.want, got) {
// 				t.Log(got)
// 			}
// 		})
// 	}
// }

// func Test_matrix_Array(t *testing.T) {
// 	type fields struct {
// 		Dense *mat.Dense
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   [][]float64
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			m := &matrix{
// 				Dense: tt.fields.Dense,
// 			}
// 			if got := m.Array(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("matrix.Array() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_matrix_Search(t *testing.T) {

	tests := []struct {
		name string
		mat  core.Matrix
		x, y interface{}
		want []float64
	}{
		{
			mat:  core.New1D(1, 2, 3),
			x:    []int{0},
			y:    []int{2},
			want: []float64{3},
		},
		{
			mat: core.New2D([][]float64{
				{1, 5, 6},
				{2, 9, 2},
				{3, 3, 3},
			}),
			x:    []int{1},
			y:    []int{1},
			want: []float64{9},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			m := tt.mat
			got := m.Search(tt.x, tt.y)
			assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", got))
			if !assert.Equal(t, tt.want, got) {
				t.Log(got)
			}
		})
	}
}

func Test_matrix_ArgMax(t *testing.T) {
	tests := []struct {
		name    string
		mat     core.Matrix
		options []core.Option
		want    core.Matrix
	}{
		{
			mat:  core.New1D(1, 3, 1),
			want: core.New1D(1),
		},
		{
			mat: core.New2D([][]float64{
				{1, 3, 1},
				{9, 3, 1},
				{9, 0, 12},
			}),
			want: core.New1D(8),
		},
		{
			options: []core.Option{core.Axis(0)},
			mat: core.New2D([][]float64{
				{1, 3, 1},
				{9, 0, 12},
			}),
			want: core.New1D(1, 0, 1),
		},
		{
			options: []core.Option{core.Axis(1)},
			mat: core.New2D([][]float64{
				{1, 3, 1},
				{9, 0, 12},
			}),
			want: core.New1D(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mat.ArgMax(tt.options...)
			if !assert.Equal(t, tt.want, got) {
				t.Log(got)
			}
		})
	}
}
