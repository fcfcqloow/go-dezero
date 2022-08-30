package datasets_test

import (
	"testing"

	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/datasets"
	appif "github.com/DolkMd/go-dezero/domain/interfaces"
	"github.com/DolkMd/go-dezero/infrastructure/infragraph"
)

func TestSinCurve(t *testing.T) {
	dataset := datasets.NewSinCurve(dz.Train(true))
	graph := infragraph.New()
	x, y, labelY := []float64{}, []float64{}, []float64{}
	for i := 0; i < dataset.Len(); i++ {
		data, label := dataset.Get(i)
		x = append(x, float64(i))
		y = append(y, data.At(0, 0))
		labelY = append(labelY, label.At(0, 0))
	}
	p1, _ := graph.Line(x, y, appif.GraphColor(255, 0, 0, 255))
	p2, _ := graph.Line(x, labelY, appif.GraphColor(0, 255, 0, 255))

	if err := graph.SaveGraph("test.png", []appif.GraphParts{p1, p2}); err != nil {
		panic(err)
	}

}
