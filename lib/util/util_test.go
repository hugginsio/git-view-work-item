package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckErrorFalse(t *testing.T) {
	var err error
	result := CheckError(err)
	assert.False(t, result)
}
