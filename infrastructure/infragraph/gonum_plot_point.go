package infragraph

import (
	"fmt"
	"image/color"
	"math/rand"

	"gonum.org/v1/plot/plotter"

	appif "github.com/DolkMd/go-dezero/domain/interfaces"
)

type (
	pointGraph struct {
		appif.GraphParts
		value *plotter.Scatter
	}
)

func newPointGraph(xys plotter.XYs, options ...appif.Option) (appif.GraphParts, error) {
	option := appif.ApplyOption(options...)
	s, err := plotter.NewScatter(xys)
	if err != nil {
		return nil, fmt.Errorf("failed to create instance: %w", err)
	}
	if option.Color == nil {
		s.GlyphStyle.Color = color.RGBA{
			R: uint8(rand.Intn(255)),
			G: uint8(rand.Intn(255)),
			B: uint8(rand.Intn(255)),
			A: uint8(rand.Intn(255)),
		}
	} else {
		s.GlyphStyle.Color = color.RGBA{
			R: option.Color.Red,
			G: option.Color.Green,
			B: option.Color.Blue,
			A: option.Color.A,
		}
	}
	return &pointGraph{value: s}, nil
}

func (p *pointGraph) Value() interface{} {
	return p.value
}
func (*pointGraph) Type() appif.GraphType {
	return appif.PointGraph()
}
