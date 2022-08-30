package transforms

import (
	"github.com/DolkMd/go-dezero/domain/core"
	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
)

func Compose(transforms ...dz.Transform) dz.Transform {
	return func(x core.Matrix) core.Matrix {
		for _, tf := range transforms {
			x = tf(x)
		}
		return x
	}
}

func Normalize(mean float64, std float64) dz.Transform {
	if std == 0 {
		std = 1
	}
	return func(x core.Matrix) core.Matrix {
		x = x.CopySubFloat(mean)
		return x.CopyDivFloat(std)
	}
}
