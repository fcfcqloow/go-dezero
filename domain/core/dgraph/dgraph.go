package dgraph

import (
	"time"

	cnv "github.com/fcfcqloow/go-advance/convert"
)

type DGraphOption func(*option)
type option struct {
	filename string
}

func ApplyDGraphOption(options ...DGraphOption) (option option) {
	for _, opt := range options {
		opt(&option)
	}

	if option.filename == "" {
		option.filename = "./graph/" + cnv.MustStr(time.Now().UnixNano()) + "_dgraph.png"
	}

	return
}

func FileName(filename string) DGraphOption {
	return func(o *option) {
		o.filename = filename
	}
}
