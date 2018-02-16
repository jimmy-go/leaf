package leaf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFloat(t *testing.T) {
	table := []struct {
		Purpose, In string
		Exp         bool
	}{
		{"1. OK", "1.23", true},
		{"2. Fail: letter", "a1.23", false},
		{"3. Fail: empty", "", false},
	}
	for _, x := range table {
		actual := IsFloat(x.In)
		assert.EqualValues(t, x.Exp, actual, x.Purpose)
	}
}

func TestIsInteger(t *testing.T) {
	table := []struct {
		Purpose, In string
		Exp         bool
	}{
		{"1. OK", "1", true},
		{"2. Fail: letter", "a1", false},
		{"3. Fail: empty", "", false},
	}
	for _, x := range table {
		actual := IsInteger(x.In)
		assert.EqualValues(t, x.Exp, actual, x.Purpose)
	}
}

func TestIsLetter(t *testing.T) {
	table := []struct {
		Purpose, In string
		Exp         bool
	}{
		{"1. OK", "a", true},
		{"2. Fail: symbol", "{", false},
		{"3. Fail: empty", "", false},
	}
	for _, x := range table {
		actual := IsLetter(x.In)
		assert.EqualValues(t, x.Exp, actual, x.Purpose)
	}
}
