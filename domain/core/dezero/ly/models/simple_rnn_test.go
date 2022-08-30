package models_test

import (
	"testing"

	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/DolkMd/go-dezero/domain/core/dezero/ly/models"
)

func TestSimpleRNN(t *testing.T) {
	seqData := dz.NewVariables()
	for i := 0; i < 1000; i++ {
		seqData = append(seqData, dz.NewVariable(core.NewRandN(core.Shape{R: 1, C: 1})))
	}

	xs := seqData[0 : len(seqData)-1]
	ts := seqData[1:]
	model := models.NewSimpleRNN(10, 1)

	lossValue, cnt := dz.NewVariable(core.NewEmpty()), 0
	for i := 0; i < len(xs); i++ {
		x := xs[i]
		t := ts[i]
		y := model.Apply(x).First()

		loss := fn.MeanSquaredError(y, t)
		if lossValue.Data().IsEmpty() {
			lossValue = loss
		} else {
			lossValue = fn.Add(lossValue, loss)
		}

		cnt += 1
		if cnt == 2 {
			model.ClearGrads()
			lossValue.Backward()
			break
		}
	}
	// fmt.Println(model.Plot(xs, models.FileName("test")))
}
