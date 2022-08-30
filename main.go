package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/DolkMd/go-dezero/infrastructure/infragraph"
	"github.com/DolkMd/go-dezero/infrastructure/sysutil"
	"github.com/cheggaaa/pb/v3"
	"github.com/fcfcqloow/go-advance/log"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/datasets"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/DolkMd/go-dezero/domain/core/dezero/loader"
	models "github.com/DolkMd/go-dezero/domain/core/dezero/ly/models"
	"github.com/DolkMd/go-dezero/domain/core/dezero/optimizers"
	"github.com/DolkMd/go-dezero/domain/core/dgraph"
	appif "github.com/DolkMd/go-dezero/domain/interfaces"
)

func main() {
	// log.SetLevel(log.LOG_LEVEL_DEBUG)
	log.SetFunctionSkip(log.DEFAULT_SKIP + 1)
	log.DebugOutput = func(values []interface{}, filePath log.FilePath, line log.ProgramLine, name log.MethodName) []string {
		return []string{"[", string(filePath), "]: ", strings.Join(log.InterfacesToStrings(values), " ")}
	}
	sysutil.RunViewMemoryLocal()
	sampleLSTM()
}

func sampleMLP() {
	const (
		lr         = 0.2
		maxIter    = 10000
		hiddenSize = 10
		fileName   = "test.w"
	)

	x := dz.NewVariable(core.NewRand(core.Shape{R: 100, C: 1}))
	a1 := x.Data().CopyApply(func(f float64) float64 { return math.Sin(2 * math.Pi * f) })
	a2 := core.NewRand(core.Shape{R: 100, C: 1})
	y := fn.Add(dz.NewVariable(a1), dz.NewVariable(a2))

	model := models.NewMLP([]int{hiddenSize, 2})
	optim := optimizers.NewSGD(dz.Lr(lr)).Setup(model)

	if err := model.LoadWeights(fileName); err != nil {
		log.Warn(err)
	}

	for i := 0; i < maxIter; i++ {
		yPred := model.Apply(x).First()
		loss := fn.MeanSquaredError(y, yPred)

		model.ClearGrads()
		loss.Backward(dz.RetainGrad(true))
		optim.Update()

		if i%100 == 0 {
			fmt.Println(loss)
		}
	}

	if err := model.SaveWeights(fileName); err != nil {
		panic(err)
	}
}
func sampleSpiral() {
	const (
		batchSize  = 30
		maxEpoch   = 300
		hiddenSize = 10
		lr         = 1.0
	)

	trainSet := datasets.NewSpiral(dz.Train(true))
	testSet := datasets.NewSpiral(dz.Train(false))
	trainLoader := dz.NewDataLoader(trainSet, batchSize)
	testloader := dz.NewDataLoader(testSet, batchSize, dz.DShuffle(true))
	model := models.NewMLP([]int{hiddenSize, 3})
	optim := optimizers.NewSGD(dz.Lr(lr)).Setup(model)
	for i := 0; i < maxEpoch; i++ {
		sumLoss, sumAcc := 0., 0.
		bar := pb.StartNew(trainLoader.Len())
		for trainLoader.Next() {
			bar.Increment()
			x, t := trainLoader.Read()
			y := model.Apply(x).First()
			loss := fn.SoftmaxCrossEntropy(y, t)
			acc := fn.Accuacy(y, t)
			model.ClearGrads()
			loss.Backward(dz.RetainGrad(true))
			optim.Update()

			sumLoss += loss.Data().At(0, 0) * float64(t.Data().Len())
			sumAcc += acc.Data().At(0, 0) * float64(t.Data().Len())
		}
		bar.Finish()
		fmt.Println("epoch", i+1, " train: ",
			"loss", sumLoss/float64(trainSet.Len()),
			"accuracy", sumAcc/float64(trainSet.Len()))

		sumLoss, sumAcc = 0, 0
		core.NoGrad(func() error {
			for testloader.Next() {
				x, t := testloader.Read()
				y := model.Apply(x).First()
				loss := fn.SoftmaxCrossEntropy(y, t)
				acc := fn.Accuacy(y, t)
				sumLoss += loss.Data().At(0, 0) * float64(t.Data().Len())
				sumAcc += acc.Data().At(0, 0) * float64(t.Data().Len())
			}
			fmt.Println("epoch", i+1, " test: ",
				"loss", sumLoss/float64(testSet.Len()),
				"accuracy", sumAcc/float64(testSet.Len()))
			return nil
		})

	}
	Contour(model)
}
func sampleMnist() {
	const (
		maxEpoch   = 5
		batchSize  = 100
		hiddenSize = 1000
	)

	lossAcc := dgraph.NewLossAcc("train", infragraph.New())
	testLossAcc := dgraph.NewLossAcc("test", infragraph.New())
	gray := func(m core.Matrix) core.Matrix { return m.CopyDivFloat(255) }
	trainSet := datasets.NewMnist(dz.Train(true), dz.TransformData(gray))
	testSet := datasets.NewMnist(dz.Train(false), dz.TransformData(gray))
	trainLoader := dz.NewDataLoader(trainSet, batchSize)
	testloader := dz.NewDataLoader(testSet, batchSize, dz.DShuffle(false))

	model := models.NewMLP([]int{hiddenSize, hiddenSize, 10}, models.ActivationFunc(func(v ...dz.Variable) dz.Variables { return fn.NewRelu().Apply(v...) }))
	optim := optimizers.NewAdam().Setup(model)
	for i := 0; i < maxEpoch; i++ {
		sumLoss, sumAcc := 0., 0.
		bar := pb.StartNew(trainLoader.Len())
		for trainLoader.Next() {
			bar.Increment()
			x, t := trainLoader.Read()
			model.Plot([]dz.Variable{x})

			y := model.Apply(x).First()
			loss := fn.SoftmaxCrossEntropy(y, t)
			acc := fn.Accuacy(y, t)
			model.ClearGrads()
			loss.Backward(dz.RetainGrad(true))
			optim.Update()

			sumLoss += loss.Data().At(0, 0) * float64(t.Data().Len())
			sumAcc += acc.Data().At(0, 0) * float64(t.Data().Len())
		}
		bar.Finish()
		lossAcc.Add(sumLoss/float64(trainSet.Len()), sumAcc/float64(trainSet.Len()))
		fmt.Println("epoch", i+1, " train: ",
			"loss", sumLoss/float64(trainSet.Len()),
			"accuracy", sumAcc/float64(trainSet.Len()))
		sumLoss, sumAcc = 0, 0
		core.NoGrad(func() error {
			for testloader.Next() {
				x, t := testloader.Read()
				y := model.Apply(x).First()
				loss := fn.SoftmaxCrossEntropy(y, t)
				acc := fn.Accuacy(y, t)
				sumLoss += loss.Data().At(0, 0) * float64(t.Data().Len())
				sumAcc += acc.Data().At(0, 0) * float64(t.Data().Len())
			}
			testLossAcc.Add(sumLoss/float64(testSet.Len()), sumAcc/float64(testSet.Len()))

			fmt.Println("epoch", i+1, " test: ",
				"loss", sumLoss/float64(testSet.Len()),
				"accuracy", sumAcc/float64(testSet.Len()))
			return nil
		})
	}

	if err := lossAcc.Plot(); err != nil {
		panic(err)
	}
	if err := testLossAcc.Plot(); err != nil {
		panic(err)
	}
}
func sampleRNN() {
	const (
		maxEpoch   = 100
		hiddenSize = 100
		bpttLength = 30
	)

	trainSet := datasets.NewSinCurve(dz.Train(true))
	seqlen := trainSet.Len()

	model := models.NewSimpleRNN(hiddenSize, 1)
	optimizer := optimizers.NewAdam().Setup(model)

	for i := 0; i < maxEpoch; i++ {
		model.ResetState()

		loss, count := dz.NewVariable(core.New1D(0)), 0
		for j := 0; j < trainSet.Len(); j++ {
			xData, tData := trainSet.Get(j)
			x := dz.AsVariable(xData)
			t := dz.AsVariable(tData)
			y := model.Apply(x).First()
			lossVar := fn.MeanSquaredError(y, t)
			loss = fn.Add(loss, lossVar)

			if count += 1; (count%bpttLength) == 0 || count == seqlen {
				model.ClearGrads()
				loss.Backward()
				loss.UnchainBackward()
				optimizer.Update()
			}
		}
		avgLoss := loss.Data().At(0, 0) / float64(count)
		fmt.Println("epoch", i+1, " | loss", avgLoss)
	}

	if err := model.SaveWeights("./weights/rnn.w"); err != nil {
		panic(err)
	}

	model.ResetState()
	x, _ := trainSet.Get(0)
	model.Plot([]dz.Variable{dz.AsVariable(x)})
	graph := infragraph.New()
	xs := core.Linspace(0, 4*math.Pi, 1000).CopyApply(math.Cos)
	predList := []float64{}
	no := []float64{}
	xValue := []float64{}
	core.NoGrad(func() error {
		for i := 0; i < xs.Len(); i++ {
			x := dz.AsVariable(core.New1D(xs.At(0, i)))
			y := model.Apply(x).First()
			no = append(no, float64(i))
			predList = append(predList, y.Data().At(0, 0))
			xValue = append(xValue, x.Data().At(0, 0))
		}
		return nil
	})

	p1, _ := graph.Line(no, predList, appif.Bule(255))
	p2, _ := graph.Line(no, xValue, appif.Green(255))
	if err := graph.SaveGraph("test.png", []appif.GraphParts{p1, p2}); err != nil {
		panic(err)
	}
}
func sampleLSTM() {
	const (
		maxEpoch   = 100
		hiddenSize = 100
		batchSize  = 30
		bpttLength = 30
	)

	trainSet := datasets.NewSinCurve(dz.Train(true))
	dataloader := loader.NewSeqDataLoader(trainSet, batchSize)
	seqlen := trainSet.Len()

	model := models.NewBetterRNN(hiddenSize, 1)
	optimizer := optimizers.NewAdam().Setup(model)

	for i := 0; i < maxEpoch; i++ {
		model.ResetState()

		loss, count := dz.NewVariable(core.New1D(0)), 0
		for dataloader.Next() {
			x, t := dataloader.Read()
			y := model.Apply(x).First()
			lossVar := fn.MeanSquaredError(y, t)
			loss = fn.Add(loss, lossVar)
			if count += 1; (count%bpttLength) == 0 || count == seqlen {
				model.ClearGrads()
				loss.Backward()
				loss.UnchainBackward()
				optimizer.Update()
			}
		}
		avgLoss := loss.Data().At(0, 0) / float64(count)
		fmt.Println("epoch", i+1, " | loss", avgLoss)
	}

	if err := model.SaveWeights("./weights/rnn.w"); err != nil {
		panic(err)
	}

	model.ResetState()
	x, _ := trainSet.Get(0)
	model.Plot([]dz.Variable{dz.AsVariable(x)})
	graph := infragraph.New()
	xs := core.Linspace(0, 4*math.Pi, 1000).CopyApply(math.Cos)
	predList := []float64{}
	no := []float64{}
	xValue := []float64{}
	core.NoGrad(func() error {
		for i := 0; i < xs.Len(); i++ {
			x := dz.AsVariable(core.New1D(xs.At(0, i)))
			y := model.Apply(x).First()
			no = append(no, float64(i))
			predList = append(predList, y.Data().At(0, 0))
			xValue = append(xValue, x.Data().At(0, 0))
		}
		return nil
	})

	p1, _ := graph.Line(no, predList, appif.Bule(255))
	p2, _ := graph.Line(no, xValue, appif.Green(255))
	if err := graph.SaveGraph("test2.png", []appif.GraphParts{p1, p2}); err != nil {
		panic(err)
	}
}

func Contour(model models.Model) {

	result := [][]float64{}
	for i := -1.0; i < 1; i += 0.01 {
		for j := -1.0; j < 1; j += 0.01 {
			result = append(result, []float64{i, j})
		}
	}
	m := dz.NewVariable(core.New2D(result))
	pred := model.Apply(m).First().Data().ArgMax(core.Axis(1))
	result = [][]float64{}
	for i := 0; i < 200; i++ {
		result = append(result, []float64{})
		for j := 0; j < 200; j++ {
			result[i] = append(result[i], pred.At(0, i*200+j))
		}
	}

	var (
		grid   = unitGrid{core.New2D(result)}
		levels = []float64{}
		c      = plotter.NewContour(
			grid,
			levels,
			palette.Rainbow(10, palette.Blue, palette.Red, 1, 1, 1),
		)
	)

	p := plot.New()
	p.Title.Text = "Contour"

	p.Add(c)

	err := p.Save(10*vg.Centimeter, 10*vg.Centimeter, "graph/contour.png")
	if err != nil {
		panic(err)
	}
}

type unitGrid struct{ mat.Matrix }

func (g unitGrid) Dims() (c, r int) { return g.Matrix.Dims() }
func (g unitGrid) Z(c, r int) float64 {
	return g.Matrix.At(c, r)
}
func (g unitGrid) X(c int) float64 {
	_, n := g.Matrix.Dims()
	if c < 0 || c >= n {
		panic("index out of range")
	}
	return float64(c)
}
func (g unitGrid) Y(r int) float64 {
	m, _ := g.Matrix.Dims()
	if r < 0 || r >= m {
		panic("index out of range")
	}
	return float64(r)
}
