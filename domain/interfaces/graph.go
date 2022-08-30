package appif

import (
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type Graph interface {
	SaveDenseImage(v dz.Variable, fileName string) error
	SaveDense(v dz.Variable, fileName string) error
	SaveGraph(filename string, graphs []GraphParts) error
	Points(x, y []float64, option ...Option) (GraphParts, error)
	Function(fn func(x float64) float64) (GraphParts, error)
	Line(x, y []float64, option ...Option) (GraphParts, error)
}

type (
	GraphParts interface {
		Value() interface{}
		Type() GraphType
	}
	GraphType struct{ graphType string }
)

func LineGraph() GraphType {
	return GraphType{graphType: "line"}
}
func PointGraph() GraphType {
	return GraphType{graphType: "point"}
}
func FunctionGraph() GraphType {
	return GraphType{graphType: "function"}
}
