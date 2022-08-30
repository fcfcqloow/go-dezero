package dz_test

import (
	"testing"

	dz "github.com/DolkMd/go-dezero/domain/core/dezero"
	"github.com/DolkMd/go-dezero/domain/core/dezero/fn"
	"github.com/stretchr/testify/assert"
)

func TestDamy(t *testing.T) {
	add := dz.NewAdd().Apply
	addfn := fn.Add
	a1 := dz.AsVariable([]float64{1, 2, 3})
	b1 := dz.AsVariable([][]float64{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}})
	a2 := dz.AsVariable([]float64{1, 2, 3})
	b2 := dz.AsVariable([][]float64{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}})
	res1 := add(a1, b1).F()
	res2 := addfn(a2, b2)
	res1.Backward()
	res2.Backward()
	assert.Equal(t, res1.Data(), res2.Data())
	assert.Equal(t, a1.Grad().Data(), a2.Grad().Data())
}
