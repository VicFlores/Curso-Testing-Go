package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddSuccess(t *testing.T) {
	c := require.New(t)

	result := Add(20, 6)

	expected := 26

	c.Equal(expected, result)

}
