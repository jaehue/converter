package converter_test

import (
	"testing"

	"github.com/jaehue/converter"
	"github.com/stretchr/testify/assert"
)

func TestConvertStringToInt(t *testing.T) {
	i := converter.StringToInt64("250万(元)")
	assert.EqualValues(t, i, 250)

	i = converter.StringToInt64("50%")
	assert.EqualValues(t, i, 50)
}

func TestConvertStringToFloat(t *testing.T) {
	i := converter.StringToFloat64("250万(元)")
	assert.EqualValues(t, i, float64(250))

	f := converter.StringToFloat64("250.50万(元)")
	assert.EqualValues(t, f, 250.50)
}
