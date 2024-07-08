package config

import (
	"log/slog"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Environment struct {
	NAME              string
	ENV               env
	VERSION           string
	BINARY_NAME       string
	DOCKER_IMAGE_NAME string
	DIST_DIR          string
	BUILD_TAGS        string
	PORT              string
	LOG_LEVEL         slog.Level
}

type env int

const (
	Dev env = iota
	Staging
	Prod
)

func (e env) String() string {
	switch e {
	case Dev:
		return "dev"
	case Staging:
		return "staging"
	case Prod:
		return "prod"
	default:
		return "dev"
	}
}

func loadEnvironment() *Environment {
	if err := godotenv.Load(); err != nil {
		return defaultEnvironment()
	}
	return buildEnvironment()
}

func buildEnvironment() *Environment {
	var envVars = map[string]func(*Environment, string){
		"APP_NAME":          func(e *Environment, v string) { e.NAME = v },
		"ENVIRONMENT":       func(e *Environment, v string) { ParseEnv(e, v) },
		"VERSION":           func(e *Environment, v string) { e.VERSION = v },
		"BINARY_NAME":       func(e *Environment, v string) { e.BINARY_NAME = v },
		"DOCKER_IMAGE_NAME": func(e *Environment, v string) { e.DOCKER_IMAGE_NAME = v },
		"DIST_DIR":          func(e *Environment, v string) { e.DIST_DIR = v },
		"BUILD_TAGS":        func(e *Environment, v string) { e.BUILD_TAGS = v },
		"PORT":              func(e *Environment, v string) { e.PORT = v },
		"LOG_LEVEL":         func(e *Environment, v string) { ParseLogLevel(e, v) },
	}
	return loadEnvs(envVars)
}

func ParseEnv(e *Environment, v string) {
	switch strings.ToLower(v) {
	case "dev":
		e.ENV = Dev
	case "staging":
		e.ENV = Staging
	case "prod":
		e.ENV = Prod
	default:
		e.ENV = Dev
	}
}

func ParseLogLevel(e *Environment, v string) {
	switch strings.ToLower(v) {
	case "debug":
		e.LOG_LEVEL = slog.LevelDebug
	case "info":
		e.LOG_LEVEL = slog.LevelInfo
	case "warn":
		e.LOG_LEVEL = slog.LevelWarn
	case "error":
		e.LOG_LEVEL = slog.LevelError
	default:
		e.LOG_LEVEL = defaultLogLevel
	}
}

func loadEnvs(em map[string]func(*Environment, string)) *Environment {
	e := defaultEnvironment()
	for k, f := range em {
		if v := os.Getenv(k); v != "" {
			f(e, v)
		}
	}
	return e
}

const defaultLogLevel = slog.LevelInfo

func defaultEnvironment() *Environment {
	return &Environment{
		NAME:              "go-base-app",
		ENV:               Dev,
		VERSION:           "v0.0.1",
		BINARY_NAME:       "gobase",
		DOCKER_IMAGE_NAME: "ghcr.io/zarldev/go-base",
		DIST_DIR:          "dist",
		BUILD_TAGS:        "api,ui",
		PORT:              ":8080",
		LOG_LEVEL:         defaultLogLevel,
	}
}
