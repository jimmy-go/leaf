package leaf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	le, err := Open("")
	assert.Nil(t, err)
	assert.NotNil(t, le)
}
