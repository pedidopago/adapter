package atob_test

import (
	"testing"

	"github.com/pedidopago/adapter/example/atob"
	"github.com/stretchr/testify/assert"
)

func TestAToB(t *testing.T) {
	tests := []struct {
		Input  atob.UserA
		Expect atob.UserB
	}{
		{
			Input: atob.UserA{
				Name:  "John",
				Age:   30,
				Score: 10.5,
				Props: []atob.PropA{
					{Name: "prop1", Value: "1"},
					{Name: "prop2", Value: "2"},
				},
			},
			Expect: atob.UserB{
				Name:     "John",
				Age:      30,
				MaxScore: "10.5",
				Props: []atob.PropB{
					{Name: "prop1", Value: 1},
					{Name: "prop2", Value: 2},
				},
			},
		},
		{
			Input: atob.UserA{
				Name:  "John Doe",
				Age:   127,
				Score: 8,
				Props: []atob.PropA{
					{Name: "prop4", Value: "1367"},
					{Name: "prop1", Value: "332"},
					{Name: "propx", Value: "banana"},
				},
			},
			Expect: atob.UserB{
				Name:     "John Doe",
				Age:      127,
				MaxScore: "8",
				Props: []atob.PropB{
					{Name: "prop4", Value: 1367},
					{Name: "prop1", Value: 332},
					{Name: "propx", Value: 0},
				},
			},
		},
	}
	for i, test := range tests {
		result := atob.UserAtoB(test.Input)
		assert.Equal(t, test.Expect.Name, result.Name, "test %d name", i)
		assert.Equal(t, test.Expect.MaxScore, result.MaxScore, "test %d name", i)
		if assert.Equal(t, len(test.Expect.Props), len(result.Props), "test %d props", i) {
			for j := range test.Expect.Props {
				assert.Equal(t, test.Expect.Props[j].Name, result.Props[j].Name, "test %d prop %d name", i, j)
				assert.Equal(t, test.Expect.Props[j].Value, result.Props[j].Value, "test %d prop %d value", i, j)
			}
		}
		assert.Equal(t, test.Expect.MaxScore, result.MaxScore, "test %d maxscore", i)
	}
}
