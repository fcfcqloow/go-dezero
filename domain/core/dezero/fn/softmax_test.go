package fn_test

// func TestSoftmax(t *testing.T) {
//
// 	testCases := map[string]struct {
// 		calc   func() []core.Matrix
// 		result []core.Matrix
// 	}{
// 		"success: test Square model": {
// 			calc: func() []core.Matrix {
// 				x := dz.NewVariable(core.New1D(2.0))
// 				y := fn.Softmax(x)
// 				return []core.Matrix{y.Data()}
// 			},
// 			result: []core.Matrix{core.New1D(4.0)},
// 		},
// 	}

// 	for name, tc := range testCases {
// 		tc := tc
// 		name := name
// 		t.Run(name, func(t *testing.T) {
//
// 			assert.Equal(t, tc.result, tc.calc())
// 		})
// 	}

// }
