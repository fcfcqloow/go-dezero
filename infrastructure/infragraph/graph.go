package infragraph

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/fcfcqloow/go-advance/log"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	apperr "github.com/DolkMd/go-dezero/domain/errors"
	appif "github.com/DolkMd/go-dezero/domain/interfaces"
)

const (
	dotTopFormat  = "digraph g {\n%s}"
	dotVarFortmat = "  \"%p\" [label=\"%s\", color=orange, style=filled]\n"
	dotFuncFormat = "  \"%p\" [label=\"%s\", color=lightblue, style=filled, shape=box]\n"
	dotEdgeFormat = "  \"%p\" -> \"%p\"\n"
)

type graph struct{}

func New() appif.Graph {
	return &graph{}
}

func (g *graph) SaveGraph(filename string, graphs []appif.GraphParts) error {
	p := plot.New()
	for _, graph := range graphs {
		log.Debug(graph.Type())
		switch graph.Type() {
		case appif.LineGraph():
			line, ok := graph.Value().(*plotter.Line)
			if !ok {
				return fmt.Errorf("missing type point graph")
			}
			p.Add(line)
		case appif.PointGraph():
			scatter, ok := graph.Value().(*plotter.Scatter)
			if !ok {
				return fmt.Errorf("missing type point graph")
			}
			p.Add(scatter)
		case appif.FunctionGraph():
			function, ok := graph.Value().(*plotter.Function)
			if !ok {
				return fmt.Errorf("missing type function graph")
			}
			p.Add(function)
		default:
			return fmt.Errorf("unknown graph type")
		}
	}
	log.Debug("Save a graph: ", filename)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, filename); err != nil {
		return fmt.Errorf("graph save error: %w", err)
	}

	return nil
}

func (g *graph) Function(fn func(x float64) float64) (appif.GraphParts, error) {
	return newFunctionGraph(fn), nil
}

func (g *graph) Points(x, y []float64, option ...appif.Option) (appif.GraphParts, error) {
	if len(x) != len(y) {
		return nil, apperr.NewPointGraphErr(fmt.Errorf("failed to crete point graph"))
	}

	xys := plotter.XYs{}
	for i := 0; i < len(x); i++ {
		xys = append(xys, plotter.XY{X: x[i], Y: y[i]})
	}

	gp, err := newPointGraph(xys, option...)
	if err != nil {
		return nil, apperr.NewPointGraphErr(err)
	}

	return gp, nil
}

func (g *graph) Line(x, y []float64, option ...appif.Option) (appif.GraphParts, error) {
	if len(x) != len(y) {
		return nil, apperr.NewPointGraphErr(fmt.Errorf("failed to crete line graph"))
	}

	xys := plotter.XYs{}
	for i := 0; i < len(x); i++ {
		xys = append(xys, plotter.XY{X: x[i], Y: y[i]})
	}

	gp, err := newLineGraph(xys, option...)
	if err != nil {
		return nil, apperr.NewPointGraphErr(err)
	}

	return gp, nil
}

func (*graph) SaveDense(v dz.Variable, name string) error {
	graph := dot(v)
	log.Debug("Save dot file: ", name)
	if err := ioutil.WriteFile(name, []byte(graph), fs.ModePerm); err != nil {
		return apperr.NewSaveGraphErr(err)
	}

	return nil
}

func (*graph) SaveDenseImage(v dz.Variable, name string) error {
	graph := dot(v)
	args := strings.Split("dot -Tpng", " ")
	cmd := exec.Command(args[0], args[1:]...)
	log.Debug(args[0], args[1:])
	log.Debug(graph)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return apperr.NewSaveGraphErr(err)
	}
	io.WriteString(stdin, graph)
	stdin.Close()
	out, err := cmd.Output()
	if err != nil {
		return apperr.NewSaveGraphErr(err)
	}
	log.Debug("Save  dot image file: ", name)
	if err := ioutil.WriteFile(name, out, fs.ModePerm); err != nil {
		return apperr.NewSaveGraphErr(err)
	}
	return nil
}
