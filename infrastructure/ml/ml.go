package ml

// import (
// 	"math/rand"
// 	"time"

// 	"github.com/DolkMd/go-dezero/domain/core"
// 	"github.com/DolkMd/go-dezero/domain/core"
// 	. "github.com/DolkMd/go-dezero/domain/fn"
// 	"github.com/DolkMd/go-dezero/domain/fn/fnlayer"
// 	"github.com/DolkMd/go-dezero/domain/util"
// )

// func Script1() {
// 	A := fnlayer.NewSquare()
// 	B := fnlayer.NewExp()
// 	C := fnlayer.NewSquare()
// 	x := dz.NewVariable(core.New1D(0.5))
// 	a := A.Apply(x)
// 	b := B.Apply(a...)
// 	y := C.Apply(b...)
// 	util.PrintDences(y.DataArr()...)
// }

// // func Script2() {
// // 	f := fnlayer.NewSquare()
// // 	x := dz.NewVariable(core.New1D(2.0))
// // 	dy := dz.NumericalDiff(f.Apply, x, nil)
// // 	util.PrintDence(dy)

// // }

// // func Script3() {
// // 	x := dz.NewVariable(core.New1D(0.5))
// // 	dy := dz.NumericalDiff(fnlayer.Composite(
// // 		fnlayer.NewSquare(),
// // 		fnlayer.NewExp(),
// // 		fnlayer.NewSquare(),
// // 	), x, nil)
// // 	util.PrintDence(dy)
// // }

// // func Script4() {
// // 	A := fnlayer.NewSquare()
// // 	B := fnlayer.NewExp()
// // 	C := fnlayer.NewSquare()
// // 	x := dz.NewVariable(core.New1D(0.5))
// // 	a := A.Apply(x)
// // 	b := B.Apply(a)
// // 	y := C.Apply(b)
// // 	util.PrintDence(y.Data()) // 1.648721270700128
// // 	y.SetGrad(core.New1D(1.0))
// // 	b.SetGrad(C.Backward(y.Grad()))
// // 	a.SetGrad(B.Backward(b.Grad()))
// // 	x.SetGrad(A.Backward(a.Grad()))
// // 	util.PrintDence(x.Grad()) // 3.297442541400256
// // }

// // func Script5() {
// // 	A := fnlayer.NewSquare()
// // 	B := fnlayer.NewExp()
// // 	C := fnlayer.NewSquare()
// // 	x := dz.NewVariable(core.New1D(0.5))
// // 	a := A.Apply(x)
// // 	b := B.Apply(a)
// // 	y := C.Apply(b)
// // 	util.PrintDence(y.Data()) // 1.648721270700128
// // 	y.SetGrad(core.New1D(1.0))
// // 	Cfn := y.Creator()
// // 	b = Cfn.Input()
// // 	b.SetGrad(Cfn.Backward(y.Grad()))
// // 	Bfn := b.Creator()
// // 	a = Bfn.Input()
// // 	a.SetGrad(Bfn.Backward(b.Grad()))
// // 	Afn := a.Creator()
// // 	x = Afn.Input()
// // 	x.SetGrad(Afn.Backward(a.Grad()))
// // 	util.PrintDence(x.Grad()) // 3.297442541400256
// // }

// func Script6() {
// 	A := fnlayer.NewSquare()
// 	B := fnlayer.NewExp()
// 	C := fnlayer.NewSquare()
// 	x := dz.NewVariable(core.New1D(0.5))
// 	a := A.Apply(x)
// 	b := B.Apply(a...)
// 	y := C.Apply(b...)
// 	util.PrintDence(y.First().Data()) // 1.648721270700128
// 	y.First().SetGrad(core.New1D(1.0))
// 	y.First().Backward()
// 	util.PrintDence(x.Grad()) // 3.297442541400256
// }

// // func Script7() {
// // 	x := dz.NewVariable(core.New1D(0.5))
// // 	a := Square(x)
// // 	b := Exp(a)
// // 	y := Square(b)
// // 	y.Backward()
// // 	util.PrintDence(x.Grad()) // 3.297442541400256
// // }

// func Script8() {
// 	x1, x2 := core.New1D(2), core.New1D(3)
// 	y := Add(dz.NewVariables(x1, x2)...)
// 	util.PrintDence(y.First().Data())
// }

// func Script9() {
// 	x := dz.NewVariable(core.New1D(2.0))
// 	y := dz.NewVariable(core.New1D(3.0))
// 	z := Add(Square(x).First(), Square(y).First())
// 	z.First().Backward()

// 	util.PrintDences(z.First().Data())
// 	util.PrintDences(x.Grad())
// 	util.PrintDences(y.Grad())
// }

// func Script10() {
// 	x := dz.NewVariable(core.New1D(3.0))
// 	y := Add(x, x)
// 	y.First().Backward()
// 	util.PrintDences(x.Grad()) // 2

// 	x.ClearGrad()
// 	y = Add(Add(x, x).First(), x)
// 	y.First().Backward()
// 	util.PrintDence(x.Grad()) // 3
// }

// func Script11() {
// 	x := dz.NewVariable(core.New1D(2.0))       // x = 2
// 	a := Square(x).First()                         // x^2 = 4
// 	y := Add(Square(a).First(), Square(a).First()) // x^2^2 + x^2^2 = 2(x^4)= 32
// 	y.First().Backward()                           // y' = 8x^3 = 64

// 	util.PrintDence(y.First().Data())
// 	util.PrintDence(x.Grad())
// }

// func random(min, max float64) float64 {
// 	rand.Seed(time.Now().UnixNano())
// 	return rand.Float64()*(max-min) + min
// }
// func Script12() {
// 	list := []float64{}
// 	for i := 0; i < 10000; i++ {
// 		list = append(list, random(0, 10000))
// 	}

// 	for i := 0; i < 10; i++ {
// 		x := dz.NewVariable(core.New1D(list...))
// 		y := Composite(Square, Square, Square)(x)
// 		util.PrintDences(y.ToDataList()...)

// 	}
// }

// func Script13() {
// 	a := dz.NewVariable(core.New1D(3.0))
// 	b := dz.NewVariable(core.New1D(2.0))
// 	c := dz.NewVariable(core.New1D(1.0))
// 	y := Add(Mul(a, b).First(), c)
// 	y.First().Backward()
// 	util.PrintDences(y.First().Data(), a.Grad(), b.Grad())

// 	x := dz.NewVariable(core.New1D(2.0))
// 	y = Neg(x)
// 	util.PrintDence(y.First().Data())

// 	x = dz.NewVariable(core.New1D(2.0))
// 	y1 := Sub(dz.NewVariable(core.New1D(2.0)), x)
// 	y2 := Sub(x, dz.NewVariable(core.New1D(1.0)))
// 	util.PrintDences(y1.First().Data(), y2.First().Data())
// }

// func MySin(x dz.Variable, threshold float64) dz.Variable {
// 	if threshold == 0 {
// 		threshold = 0.0001
// 	}
// 	y := dz.NewVariable(core.New1D(0))
// 	for i := 0; i < 100000; i++ {
// 		tmp := math.Pow(-1, float64(i))
// 		c := big.NewFloat(1).Quo(big.NewFloat(tmp), factorial(2*i+1))
// 		t := fn.MulFloat(fn.GenPowFunc(float64(2*i + 1))(x)[0], asFloat64(c))
// 		y = fn.Add(y, t[0]).First()
// 		if math.Abs(t.F().Data().At(0, 0)) < threshold {
// 			break
// 		}
// 	}

// 	return y

// }
// func asFloat64(f *big.Float) float64 {
// 	r, _ := f.Float64()
// 	return r
// }

// func factorial(n interface{}) *big.Float {
// 	return new(big.Float).SetInt(big.NewInt(1).MulRange(1, cnv.MustInt64(n)))
// }
