package config

import (
	"log"
	"os"
)

type Environment struct {
	Port            string
	Secret          string
	MongoUrl        string
	MongoDb         string
	MongoCollection string
}

var Fatalf func(format string, v ...interface{}) = log.Fatalf

func OptionalEnv(key string, def string) (string) {
	env := os.Getenv(key)
	if len(env) > 0 {
		return env
	} 
	return def
}

func RequireEnv(key string) (string) {
	env := os.Getenv(key)
	if len(env) > 0 {
		return env
	} 
	Fatalf("missing env variable '%s'", key)
	return ""
}

func LoadEnv() (*Environment) {
	env := new(Environment)
	env.Port = OptionalEnv("PORT", "3000")
	env.Secret = RequireEnv("SECRET")
	env.MongoUrl = RequireEnv("MONGO_URL")
	env.MongoDb = RequireEnv("MONGO_DB")
	env.MongoCollection = RequireEnv("MONGO_COLLECTION")
	return env
}
