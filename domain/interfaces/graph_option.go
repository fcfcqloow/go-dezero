package appif

type (
	Color  struct{ Red, Green, Blue, A uint8 }
	Option func(*option)
	option struct{ Color *Color }
)

func ApplyOption(options ...Option) (result option) {
	for _, opt := range options {
		opt(&result)
	}

	return
}

func GraphColor(r, g, b, a uint8) Option {
	return func(o *option) {
		o.Color = &Color{Red: r, Green: g, Blue: b, A: a}
	}
}

func Red(a uint8) Option   { return GraphColor(255, 0, 0, a) }
func Bule(a uint8) Option  { return GraphColor(0, 0, 255, a) }
func Green(a uint8) Option { return GraphColor(0, 255, 0, a) }
