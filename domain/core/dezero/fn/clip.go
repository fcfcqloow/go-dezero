package fn

import dz "github.com/DolkMd/go-dezero/domain/core/dezero"

type (
	clip struct {
		dz.Function
		xMin float64
		xMax float64
	}
)

func NewClip(xMin, xMax float64) dz.Function {
	instance := new(clip)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "Clip")
	instance.xMin = xMin
	instance.xMax = xMax
	return instance
}

func (c *clip) Forward(variables ...dz.Variable) dz.Variables {
	x := variables[0]
	return []dz.Variable{dz.NewVariable(x.Data().CopyClip(c.xMin, c.xMax))}
}

func (c *clip) Backward(variables ...dz.Variable) dz.Variables {
	x, gy := c.Inputs()[0], variables[0]
	a1 := x.Data().OnOff(func(i, j int, v float64) bool { return v >= c.xMin })
	a2 := x.Data().OnOff(func(i, j int, v float64) bool { return v <= c.xMax })
	mask := Mul(dz.NewVariable(a1), dz.NewVariable(a2))
	gx := Mul(gy, mask)
	return []dz.Variable{gx}

}
