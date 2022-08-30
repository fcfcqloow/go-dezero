package models

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/fcfcqloow/go-advance/log"

	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

type (
	Model interface {
		dz.Layer
		Plot(xs []dz.Variable, options ...ModelOption) error
	}
	_model struct {
		dz.Layer
	}
)

func run(cmdLine string) error {
	args := strings.Split(cmdLine, " ")
	cmd := exec.Command(args[0], args[1:]...)
	return cmd.Run()
}

func ExtendsModel(layer dz.Layer) Model {
	instance := new(_model)
	instance.Layer = layer
	return instance
}

func (m *_model) Plot(xs []dz.Variable, options ...ModelOption) error {
	option := ApplyOption(options...)
	if option.toFile == "" {
		option.toFile = "./graph/" + cnv.MustStr(time.Now().UnixMilli())
	}

	y := m.Layer.Forward(xs...)

	log.Debug("Save dot file: ", option.toFile)
	if err := ioutil.WriteFile(option.toFile+".dot", []byte(GraphDenseString(y.First())), fs.ModePerm); err != nil {
		return err
	}

	if err := run(fmt.Sprintf("dot %s -T png -o %s", option.toFile+".dot", option.toFile+".png")); err != nil {
		return fmt.Errorf("failed to plot from dot file: %w", err)
	}

	if err := os.Remove(option.toFile + ".dot"); err != nil {
		log.Error(err)
		return fmt.Errorf("failed to remove dot file: %w", err)
	}

	return nil
}

type (
	ModelOption func(*modelOption)
	modelOption struct {
		toFile string
	}
)

func ApplyOption(options ...ModelOption) modelOption {
	option := modelOption{}
	for _, opt := range options {
		opt(&option)
	}

	return option
}

func FileName(toFile string) ModelOption {
	return func(mo *modelOption) {
		mo.toFile = toFile
	}
}
