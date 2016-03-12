package crc24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrc24(t *testing.T) {

	d1 := []byte{1, 2, 3, 4, 5, 6}
	assert.Equal(t, uint32(0xbc7e06), ChecksumOpenPGP(d1))

	d2 := []byte{0, 0, 0, 0}
	assert.Equal(t, uint32(0), ChecksumOpenPGP(d2))
}
