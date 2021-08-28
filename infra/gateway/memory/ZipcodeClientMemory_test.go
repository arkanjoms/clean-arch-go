package memory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnDistance(t *testing.T) {
	client := NewZipcodeClient()
	distance := client.Distance("ABC", "DEF")
	assert.Equal(t, 1000.0, distance)
}
