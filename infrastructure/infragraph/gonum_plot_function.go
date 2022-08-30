package infragraph

import (
	"image/color"
	"math/rand"

	appif "github.com/DolkMd/go-dezero/domain/interfaces"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type functionGraph struct {
	value *plotter.Function
}

func newFunctionGraph(fn func(x float64) float64, opts ...appif.Option) appif.GraphParts {
	opt := appif.ApplyOption(opts...)
	g := plotter.NewFunction(fn)
	g.Width = vg.Points(4)
	g.Dashes = []vg.Length{vg.Points(4), vg.Points(5)}
	if opt.Color == nil {
		g.Color = color.RGBA{
			R: uint8(rand.Intn(255)),
			G: uint8(rand.Intn(255)),
			B: uint8(rand.Intn(255)),
			A: uint8(rand.Intn(255)),
		}
	} else {
		g.Color = color.RGBA{
			R: opt.Color.Red,
			G: opt.Color.Green,
			B: opt.Color.Blue,
			A: opt.Color.A,
		}
	}

	return &functionGraph{value: g}
}

func (p *functionGraph) Value() interface{} {
	return p.value
}

func (*functionGraph) Type() appif.GraphType {
	return appif.FunctionGraph()
}
