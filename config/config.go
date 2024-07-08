// Description: This file is used to load the configuration for the application.
// The configuration is loaded from the environment variables and the logger is
// initialized with the default options including the log level from the
// environment variables.
//
// Environment variables are available through the ENVIRONMENT variable.
// The logger is available through the LOGGER variable or the standard slog package.
package config

var (
	ENVIRONMENT = loadEnvironment()
	LOGGER      = loadLogger()
)
