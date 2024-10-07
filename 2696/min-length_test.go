package minLength

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinLenth(t *testing.T) {
	s := "ABFCACDB"
	minLength := minLength(s)
	assert.Equal(t, 2, minLength)
}

func TestMinLenth2(t *testing.T) {
	s := "ACBBD"
	minLength := minLength(s)
	assert.Equal(t, 5, minLength)
}

func TestMinLenth3(t *testing.T) {
	s := "CCCCDDDD"
	minLength := minLength(s)
	assert.Equal(t, 0, minLength)
}
