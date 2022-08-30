package infragraph

import (
	"fmt"

	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/util"
)

func dot(v dz.Variable) string {
	fncs := []dz.Function{}
	seenSet := util.NewSet()
	addFunc := func(f dz.Function) {
		if !seenSet.Contains(f) {
			fncs = append(fncs, f)
			seenSet.Add(f)
		}
	}
	addFunc(v.Creator())
	varText := ""
	funcText := ""
	funcEdgeText := ""
	for len(fncs) > 0 {
		fn := fncs[0]
		fncs = fncs[1:]
		f, e := dotFunc(fn)
		funcText += f
		funcEdgeText += e
		for _, x := range fn.Inputs() {
			varText += dotVar(x)
			if x.Creator() != nil {
				addFunc(x.Creator())
			}
		}
	}
	varText += dotVar(v)
	return fmt.Sprintf(dotTopFormat, varText+funcText+funcEdgeText)
}
func dotVar(v dz.Variable) string {
	return fmt.Sprintf(dotVarFortmat, v, v.Name())
}
func dotFunc(fn dz.Function) (defText string, edgeText string) {
	defText = fmt.Sprintf(dotFuncFormat, fn, fn.Name())
	for _, input := range fn.Inputs() {
		edgeText += fmt.Sprintf(dotEdgeFormat, input, fn)
	}
	for _, output := range fn.Outputs() {
		edgeText += fmt.Sprintf(dotEdgeFormat, fn, output)
	}
	return defText, edgeText
}
