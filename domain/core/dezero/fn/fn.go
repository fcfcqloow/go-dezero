package fn

import (
	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

func floatCalc(x dz.Variable, f float64, fn func(x1, x2 dz.Variable) dz.Variable) dz.Variable {
	return fn(x, dz.NewVariable(core.New1D(f)))
}
func Square(x dz.Variable) dz.Variable                     { return NewSquare().Apply(x).First() }
func Exp(x dz.Variable) dz.Variable                        { return NewExp().Apply(x).First() }
func Add(x1, x2 dz.Variable) dz.Variable                   { return NewAdd().Apply(x1, x2).First() }
func Sub(x1, x2 dz.Variable) dz.Variable                   { return NewSub().Apply(x1, x2).First() }
func Mul(x1, x2 dz.Variable) dz.Variable                   { return NewMul().Apply(x1, x2).First() }
func Div(x1, x2 dz.Variable) dz.Variable                   { return NewDiv().Apply(x1, x2).First() }
func MatMul(x1, x2 dz.Variable) dz.Variable                { return NewMatMul().Apply(x1, x2).First() }
func Neg(x dz.Variable) dz.Variable                        { return NewNeg().Apply(x).First() }
func Sin(x dz.Variable) dz.Variable                        { return NewSin().Apply(x).First() }
func Cos(x dz.Variable) dz.Variable                        { return NewCos().Apply(x).First() }
func Tanh(x dz.Variable) dz.Variable                       { return NewTanh().Apply(x).First() }
func Log(x dz.Variable) dz.Variable                        { return NewLog().Apply(x).First() }
func Sum(x dz.Variable, os ...core.Option) dz.Variable     { return NewSum(os...).Apply(x).F() }
func SumTo(x dz.Variable, shape core.Shape) dz.Variable    { return NewSumTo(shape).Apply(x).F() }
func Transpose(xs ...dz.Variable) dz.Variable              { return NewTranspose().Apply(xs...).F() }
func AddFloat(x dz.Variable, f float64) dz.Variable        { return floatCalc(x, f, Add) }
func SubFloat(x dz.Variable, f float64) dz.Variable        { return floatCalc(x, f, Sub) }
func MulFloat(x dz.Variable, f float64) dz.Variable        { return floatCalc(x, f, Mul) }
func DivFloat(x dz.Variable, f float64) dz.Variable        { return floatCalc(x, f, Div) }
func Pow(x dz.Variable, c float64) dz.Variable             { return NewPow(c).Apply(x).First() }
func Linear(x, w, b dz.Variable) dz.Variable               { return NewLinear().Apply(x, w, b).First() }
func Softmax(x dz.Variable, os ...core.Option) dz.Variable { return NewSoftmax(os...).Apply(x).First() }
func Sigmoid(x dz.Variable) dz.Variable                    { return NewSigmoid().Apply(x).First() }
func Relu(x dz.Variable) dz.Variable                       { return NewRelu().Apply(x).First() }
func SoftmaxCrossEntropy(x, t dz.Variable) dz.Variable {
	return NewSoftmaxCrossEntropy().Apply(x, t)[0]
}
func Clip(x dz.Variable, xMin, xMax float64) dz.Variable {
	return NewClip(xMin, xMax).Apply(x).First()
}
func Reshape(x dz.Variable, s core.Shape) dz.Variable {
	if x.Data().Shape() == s {
		return x
	}
	return NewReshape(s).Apply(x).First()
}
func Matyas(x, y dz.Variable) dz.Variable {
	a := Add(Square(x), Square(y))                             // x^2 + y^2
	b := Mul(x, y)                                             // x*y
	c := Mul(a, dz.NewVariable(core.NewFull(a.Shape(), 0.26))) // 0.26 * (x^2 + y^2)
	d := Mul(b, dz.NewVariable(core.NewFull(a.Shape(), 0.48))) // 0.48x*y
	return Sub(c, d)                                           // 0.26 * (x^2 + y^2) - 0.48x*y
}
func Sphere(x, y dz.Variable) dz.Variable {
	a := Square(x)   // x^2
	b := Square(y)   // y^2
	return Add(a, b) // x^2 + y^2
}
func Rosenbrock(x0, x1 dz.Variable) dz.Variable {
	a1 := Square(Sub(x1, Square(x0)))
	a2 := Mul(a1, dz.NewVariable(a1.Data().CopyFull(100)))
	b1 := Square(Sub(x0, dz.NewVariable(x0.Data().CopyFull(1))))
	return Add(a2, b1)
}
func Goldstein(x, y dz.Variable) dz.Variable {
	a1 := AddFloat(Add(x, y), 1)                 // x + y + 1
	A := Square(a1)                              // (x + y + 1) ^2
	b1 := MulFloat(x, 14)                        // 14x
	b2 := MulFloat(Square(x), 3)                 // 3(x^2)
	b3 := MulFloat(y, 14)                        // 14y
	b4 := MulFloat(Mul(x, y), 6)                 // 6xy
	b5 := MulFloat(Square(y), 3)                 // 3(y^2)
	b6 := Sub(Add(Sub(Add(b5, b4), b3), b2), b1) // 3(y^2) + 6xy - 14y + 3(x^2) - 14x
	b7 := dz.NewVariable(b6.Data().CopyFull(19)) // 19
	B := Add(b6, b7)                             // 3(y^2) + 6xy - 14y + 3(x^2) - 14x + 19
	AB := AddFloat(Mul(A, B), 1)                 // 1 + (x + y + 1)^2 * (3(y^2) + 6xy - 14y + 3(x^2) - 14x + 19)

	c1 := Sub(MulFloat(x, 2), MulFloat(y, 3)) // 2x - 3y
	C := Square(c1)                           // (2x - 3y)^2
	d1 := MulFloat(x, 32)                     // 32x
	d2 := MulFloat(Square(x), 12)             // 12(x^2)
	d3 := MulFloat(y, 48)                     // 48y
	d4 := MulFloat(Mul(x, y), 36)             // 36xy
	d5 := MulFloat(Square(y), 27)             // 27(y^2)
	d6 := Add(Sub(d5, d4), d3)                // 27(y^2) -36xy + 48y
	d7 := Sub(Add(d6, d2), d1)                // 27(y^2) -36xy + 48y + 12(x^2) - 32x
	D := AddFloat(d7, 18)                     // 18 + 27(y^2) -36xy + 48y + 12(x^2) - 32x
	CD := AddFloat(Mul(C, D), 30)             // 30 + (2x - 3x)^2 * (18 + 27(y^2) -36xy + 48y + 12(x^2) - 32x))

	return Mul(AB, CD) // (1 + (x + y + 1)^2 * (19 - 14x + 3(x^2) - 14y + 6xy + 3(y^2)) ) * (30 + ((2x - 3y)^2) * (18 - 32x + 12(x^2) + 48y - 36xy + 27(y^2))))
}
func MeanSquaredError(x0, x1 dz.Variable) dz.Variable {
	return NewMeanSquaredError().Apply(x0, x1).First()
}
func LinearSimple(x, w, b dz.Variable) dz.Variable {
	t := MatMul(x, w)
	if b == nil {
		return t
	}

	y := Add(t, b)
	t.SetData(nil)

	return y
}
func SigmoidSimple(x dz.Variable) dz.Variable {
	a1 := AddFloat(Exp(Neg(x)), 1)
	a2 := dz.NewVariable(core.New1D(1))
	return Div(a2, a1)
}
func SoftmaxSimple(x dz.Variable, options ...core.Option) dz.Variable {
	y := Exp(x)
	sumY := Sum(y, options...)
	return Div(y, sumY)
}
func rangeInts(n int) []int {
	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, i)
	}
	return result
}
func SoftmaxCrossEntropySimple(x, t dz.Variable) dz.Variable {
	N := x.Shape().R
	p := SoftmaxSimple(x, core.Axis(1))
	p = Clip(p, 1e-15, 1.0)
	logP := Log(p)
	tlogP := dz.NewVariable(
		core.New1D(
			logP.Data().Search(rangeInts(N), t.Data().Flatten())...,
		),
	)
	y := MulFloat(DivFloat(Sum(tlogP), float64(N)), -1)
	return y
}
