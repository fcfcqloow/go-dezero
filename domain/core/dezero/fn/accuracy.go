package fn

import (
	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

func Accuacy(y, t dz.Variable) dz.Variable {
	pred := y.Data().ArgMax(core.Axis(1)).CopyReshape(t.Shape())
	result := pred.OnOff(func(i, j int, v float64) bool {
		return v == t.Data().At(i, j)
	})

	return dz.NewVariable(result.Mean())
}
