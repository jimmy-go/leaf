package leaf

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	wd, err := os.Getwd()
	assert.Nil(t, err)

	table := []struct {
		Purpose string
		File    string
		Exp     error
	}{
		{"1. OK: local file", wd + "/_samples/leaf.xlsx", nil},
		{"2. OK: remote file", "https://docs.google.com/spreadsheets/d/1Kd2lLGXU20l5IYEPbc7gt6D9jQ9tDIwAspvEVmlGtNQ/edit?usp=sharing", nil},
		{"3. FAIL", "notfound.xlsx", errors.New("open notfound.xlsx: no such file or directory")},
	}
	for _, x := range table {
		_, err := Open(x.File)
		assert.EqualValues(t, fmt.Sprintf("%s", x.Exp), fmt.Sprintf("%s", err), x.Purpose)
	}
}
