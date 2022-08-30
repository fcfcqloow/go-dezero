package dz

import (
	"fmt"

	"github.com/DolkMd/go-dezero/domain/util"
	cnv "github.com/fcfcqloow/go-advance/convert"
)

type Ly map[string]interface{}

type Layer interface {
	fmt.Stringer
	Params() []Parameter
	Set(name string, value interface{})
	Get(name string) interface{}
	GetLayer(name string) Layer
	GetParameter(name string) Parameter
	Apply(inputs ...Variable) Variables
	Forward(xs ...Variable) Variables
	ClearGrads()
	SaveWeights(filename string) error
	LoadWeights(filename string) error

	FlattenParams(paramDict map[string]Variable, parentKey string)
}

type layer struct {
	params          util.Set
	feilds          map[string]interface{}
	inputs, outputs Variables
	forward         func(xs ...Variable) Variables
}

func IsLayer(value interface{}) bool {
	_, isLayer := value.(Layer)
	return isLayer
}

func NewLayer() Layer {
	return &layer{
		params: util.NewSet(),
		feilds: map[string]interface{}{},
	}
}

func ExtendsLayer(forward func(...Variable) Variables) Layer {
	return &layer{
		params:  util.NewSet(),
		feilds:  map[string]interface{}{},
		forward: forward,
	}
}

func (l *layer) Apply(inputs ...Variable) Variables {
	outputs := l.forward(inputs...)
	l.inputs = inputs
	l.outputs = outputs

	return outputs
}

func (l *layer) Forward(xs ...Variable) Variables {
	return l.forward(xs...)
}

func (l *layer) Params() []Parameter {
	result := []Parameter{}
	for _, p := range l.params.Values() {
		feild := l.feilds[p.(string)]
		if IsLayer(feild) {
			result = append(result, feild.(Layer).Params()...)
		}
		if IsParameter(feild) {
			result = append(result, feild.(Parameter))
		}
	}

	return result
}

func (l *layer) ClearGrads() {
	params := l.Params()
	for i := range params {
		params[i].ClearGrad()
	}
}

func (l *layer) Set(name string, value interface{}) {
	if IsParameter(value) || IsLayer(value) {
		l.params.Add(name)
	}
	l.feilds[name] = value
}

func (l *layer) Get(name string) interface{} {
	return l.feilds[name]
}

func (l *layer) GetLayer(name string) Layer {
	return l.Get(name).(Layer)
}
func (l *layer) GetParameter(name string) Parameter {
	return l.Get(name).(Parameter)
}

func (l *layer) String() string {
	return cnv.MustStr(l.feilds)
}

func (l *layer) FlattenParams(paramDict map[string]Variable, parentKey string) {
	for _, name := range l.params.Values() {
		obj := l.Get(name.(string))
		key := name.(string)
		if parentKey != "" {
			key = parentKey + "/" + name.(string)
		}

		if lay, ok := obj.(Layer); ok {
			lay.FlattenParams(paramDict, key)
		} else {
			paramDict[key] = obj.(Parameter)
		}
	}
}

func (l *layer) SaveWeights(filename string) error {
	dict := map[string]Variable{}
	l.FlattenParams(dict, "")
	arrayDict := map[string]Variable{}
	for key, value := range dict {
		if value != nil {
			arrayDict[key] = value
		}
	}
	return SaveVariables(filename, arrayDict)
}

func (l *layer) LoadWeights(filename string) error {
	v, err := LoadVariables(filename)
	if err != nil {
		return err
	}

	dict := map[string]Variable{}
	l.FlattenParams(dict, "")
	for key, param := range dict {
		param.SetData(v[key].Data())
	}

	return nil
}
