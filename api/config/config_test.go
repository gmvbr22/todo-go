package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnv(t *testing.T) {

	err := os.Setenv("PORT", "8080")
	assert.NoError(t, err)

	err = os.Setenv("SECRET", "test")
	assert.NoError(t, err)

	env, err := LoadEnv()
	assert.NoError(t, err)
	assert.Equal(t, env.Port, "8080")
	assert.Equal(t, env.Secret, "test")
}

func TestDefaultPort(t *testing.T) {
	err := os.Setenv("PORT", "")
	assert.NoError(t, err)

	err = os.Setenv("SECRET", "test")
	assert.NoError(t, err)

	env, err := LoadEnv()
	assert.NoError(t, err)
	assert.Equal(t, env.Port, "3000")
}

func TestInvalidSecret(t *testing.T) {
	err := os.Setenv("PORT", "3000")
	assert.NoError(t, err)

	err = os.Setenv("SECRET", "")
	assert.NoError(t, err)

	env, err := LoadEnv()
	assert.Nil(t, env)
	assert.Error(t, err, INVALID_SECRET)
}
