package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddSuccess(t *testing.T) {
	assert := require.New(t)

	actual := Add(20, 2)
	expected := 22

	assert.Equal(expected, actual, "This should be equal")
	//if result != expected {
	//	t.Errorf("got %d, expected %d", result, expected)
	//}
}
