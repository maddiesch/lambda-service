package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventGreeting(t *testing.T) {
	t.Run("given a greeting and name", func(t *testing.T) {
		os.Setenv("GREETING", "Hello")

		event := Event{Name: "World"}
		greeting := greeting(event)

		assert.Equal(t, greeting, "Hello World!")
	})
}
