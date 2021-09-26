package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionalEnv(t *testing.T) {

	err := os.Setenv("PORT", "8080")
	assert.NoError(t, err)
	assert.Equal(t, "8080", OptionalEnv("PORT", "8080"))

	err = os.Setenv("PORT", "")
	assert.NoError(t, err)
	assert.Equal(t, "3000", OptionalEnv("PORT", "3000"))
}

func TestRequireEnv(t *testing.T) {

	Fatalf = func(format string, v ...interface{}) {
		assert.Equal(t, "missing env variable 'PORT'", fmt.Sprintf(format, v...))
	}

	err := os.Setenv("PORT", "8080")
	assert.NoError(t, err)
	assert.Equal(t, "8080", RequireEnv("PORT"))

	err = os.Setenv("PORT", "")
	assert.NoError(t, err)

	RequireEnv("PORT")
}
