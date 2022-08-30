package infragraph

import (
	"fmt"
	"image/color"
	"math/rand"

	appif "github.com/DolkMd/go-dezero/domain/interfaces"
	"gonum.org/v1/plot/plotter"
)

type (
	LineGraph struct {
		appif.GraphParts
		value *plotter.Line
	}
)

func newLineGraph(xys plotter.XYs, options ...appif.Option) (appif.GraphParts, error) {
	line, err := plotter.NewLine(xys)
	if err != nil {
		return nil, fmt.Errorf("failed to new")
	}
	option := appif.ApplyOption(options...)
	if option.Color == nil {
		line.Color = color.RGBA{
			R: uint8(rand.Intn(255)),
			G: uint8(rand.Intn(255)),
			B: uint8(rand.Intn(255)),
			A: uint8(rand.Intn(255)),
		}
	} else {
		line.Color = color.RGBA{
			R: option.Color.Red,
			G: option.Color.Green,
			B: option.Color.Blue,
			A: option.Color.A,
		}
	}
	return &LineGraph{value: line}, nil

}

func (p *LineGraph) Value() interface{} {
	return p.value
}
func (*LineGraph) Type() appif.GraphType {
	return appif.LineGraph()
}
