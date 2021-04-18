package util_test

import (
	"testing"

	"github.com/alex-petrov-vt/raytracer/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestAlmostEqual(t *testing.T) {
	assert.True(t, util.AlmostEqual((1.0+2.0), 3.0))
	assert.True(t, util.AlmostEqual((1.0*2.0), 2.0))
	assert.True(t, util.AlmostEqual((1.0+2.0), 3))
	assert.True(t, util.AlmostEqual((0.05*0.05), 0.0025))
	assert.True(t, util.AlmostEqual((0.005*0.005), 0.000025))
	assert.True(t, util.AlmostEqual((1234.5*6789.0), 8381020.5))
	assert.False(t, util.AlmostEqual((1234.5*6789.0), 8381020.6))
	assert.False(t, util.AlmostEqual((1.0+2.0), 3.0001))
	assert.False(t, util.AlmostEqual((0.005*0.005), 0.000026))
}
