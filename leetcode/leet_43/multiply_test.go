package leet_43

import (
	"testing"

	"math/big"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	var testData = []struct {
		a, b string
	}{
		{
			a: "38917123124332127777738",
			b: "49172310",
		},
		{
			a: "1",
			b: "1",
		},
		{
			a: "1",
			b: "2",
		},
		{
			a: "2",
			b: "5",
		},
		{
			a: "999",
			b: "999",
		},
		{
			a: "99988",
			b: "999",
		},
		{
			a: "9991231231231203192423894893489019309201930128930912388",
			b: "0",
		},
		{
			a: "9991231231231203192423894893489019309201930128930912388",
			b: "1",
		},
		{
			a: "888",
			b: "7777777777",
		},
		{
			a: "9991231231231203192423894893489019309201930128930912388",
			b: "12938712984677182764823648236482364872364126381264387126348723648123648723647861726178",
		},
	}
	a := new(big.Int)
	b := new(big.Int)
	ass := assert.New(t)
	for _, tc := range testData {
		res := multiply(tc.a, tc.b)
		a.SetString(tc.a, 10)
		b.SetString(tc.b, 10)
		expect := big.NewInt(0).Mul(a, b)
		ass.Equal(expect.String(), res, "input a=%v b=%v", tc.a, tc.b)
	}
}
