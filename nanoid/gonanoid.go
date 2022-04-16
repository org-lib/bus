package nanoid

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// Simple usage
// Custom length

func New(l ...int) (string, error) {
	// Simple usage
	return gonanoid.New()
}

// Custom alphabet

func Generate(alphabet string, size int) (string, error) {
	return gonanoid.Generate(alphabet, size)
}

// Custom non ascii alphabet

func Must(l ...int) string {
	return gonanoid.Must()
}

// Custom non ascii alphabet

func MustGenerate(alphabet string, size int) string {
	return gonanoid.MustGenerate(alphabet, size)
}
