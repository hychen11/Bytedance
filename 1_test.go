package Bytedance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Tom() string {
	return "Tom"
}

func TestTom(t *testing.T) {
	output := Tom()
	expectOutput := "Tom"
	//if output != expectOutput {
	//	t.Errorf("%s not match %s", expectOutput, output)
	//}
	assert.Equal(t, output, expectOutput)
}
