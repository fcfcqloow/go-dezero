package dgraph

import (
	"strings"

	appif "github.com/DolkMd/go-dezero/domain/interfaces"
)

type LossAcc interface {
	Add(loss, acc float64)
	Plot() error
}
type lossAcc struct {
	loss  []float64
	acc   []float64
	graph appif.Graph
	name  string
}

func NewLossAcc(name string, graph appif.Graph) LossAcc {
	return &lossAcc{
		graph: graph,
		name:  name,
	}
}

func (l *lossAcc) Add(loss, acc float64) {
	l.loss = append(l.loss, loss)
	l.acc = append(l.acc, acc)
}

func (l *lossAcc) Plot() error {
	option := ApplyDGraphOption()
	lossX := make([]float64, len(l.loss))
	for i := range l.loss {
		lossX[i] = float64(i)
	}

	accX := make([]float64, len(l.acc))
	for i := range l.acc {
		accX[i] = float64(i)
	}

	lossPart, err := l.graph.Line(lossX, l.loss, appif.Bule(255))
	if err != nil {
		return err
	}

	accPart, err := l.graph.Line(accX, l.acc, appif.Red(255))
	if err != nil {
		return err
	}

	if err := l.graph.SaveGraph(strings.ReplaceAll(option.filename, ".png", l.name+"_loss.png"), []appif.GraphParts{lossPart}); err != nil {
		return err
	}

	return l.graph.SaveGraph(strings.ReplaceAll(option.filename, ".png", l.name+"_acc.png"), []appif.GraphParts{accPart})

}
