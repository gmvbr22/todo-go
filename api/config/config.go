package config

import (
	"errors"
	"os"
)

type Environment struct {
	Port   string
	Secret string
}

const INVALID_SECRET = "missing env variable 'secret'"

func LoadEnv() (*Environment, error) {

	env := new(Environment)
	port := os.Getenv("PORT")
	secret := os.Getenv("SECRET")

	if len(port) > 0 {
		env.Port = port
	} else {
		env.Port = "3000"
	}

	if len(secret) > 0 {
		env.Secret = secret
	} else {
		return nil, errors.New(INVALID_SECRET)
	}

	return env, nil
}
