package fn

import dz "github.com/DolkMd/go-dezero/domain/core/dezero"

type (
	matMul struct{ dz.Function }
)

func NewMatMul() dz.Function {
	instance := new(matMul)
	instance.Function = dz.ExtendsFunction(instance.Forward, instance.Backward, "MatMul")
	return instance
}

func (m *matMul) Forward(variables ...dz.Variable) dz.Variables {
	x, W := variables[0], variables[1]
	return []dz.Variable{dz.NewVariable(x.Data().CopyMatMul(W.Data()))}
}

func (m *matMul) Backward(variables ...dz.Variable) dz.Variables {
	x, W, gy := m.Inputs()[0], m.Inputs()[1], variables[0]
	gx := MatMul(gy, Transpose(W))
	gW := MatMul(Transpose(x), gy)
	return []dz.Variable{gx, gW}
}
