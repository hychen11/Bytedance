package Bytedance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Judge(70)
	assert.Equal(t, true, result)
}
