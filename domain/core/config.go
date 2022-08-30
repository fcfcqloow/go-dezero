package core

import "sync"

type (
	config struct {
		enableBackprop bool
	}
	ConfigOption func(*config)
)

var (
	Config  = config{enableBackprop: true}
	mutex   sync.Mutex
	isUsing = false
)

func (c *config) EnableBackprop() bool {
	return c.enableBackprop
}

func (c *config) SetEnableBackprop(enableBackprop bool) {
	if isUsing {
		mutex.Lock()
	}
	c.enableBackprop = enableBackprop
}

func UsingConfig(fn func() error, opts ...ConfigOption) error {
	isUsing = true
	copiedConf := Config
	defer func() {
		Config = copiedConf
		isUsing = false
		mutex.Unlock()
	}()

	for _, opt := range opts {
		opt(&Config)
	}

	return fn()
}

func NoGrad(fn func() error) error {
	return UsingConfig(fn, EnableBackprop(false))
}
func EnableBackprop(enable bool) ConfigOption {
	return func(c *config) { c.SetEnableBackprop(enable) }
}
