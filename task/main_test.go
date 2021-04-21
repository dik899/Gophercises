package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainApp(t *testing.T) {
	err := startApp()
	var f = (err == nil)

	assert.Equal(t, false, f, "Main is working ")
}

func TestMainBody(t *testing.T) {
	main()
}
