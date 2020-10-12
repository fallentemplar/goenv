package goenv

import (
	"os"
	"strconv"
)

// GetInt recovers an environment variable with the key as the received parameter
// and casts it to an int. If the requested variable doesn't exists, returns the
// default value
func GetInt(key string, defaultValue int) (value int) {
	variable := os.Getenv(key)
	if len(variable) == 0 {
		return defaultValue
	}
	value, err := strconv.Atoi(variable)
	if err != nil {
		return defaultValue
	}
	return value
}

// GetBool recovers an environment variable with the key as the received parameter
// and casts it to a boolean. If the requested variable doesn't exists, returns the
// default value
// It considers 'true' and '1' as truthy values and 'false' and '0' as falsy values
func GetBool(key string, defaultValue bool) (value bool) {
	variable := os.Getenv(key)
	if len(variable) == 0 {
		return defaultValue
	}

	if variable == "0" || variable == "false" {
		return false
	} else if variable == "1" || variable == "true" {
		return true
	}

	return defaultValue
}

// GetString recovers an environment variable with the key as the received parameter
// and casts it to an int. If the requested variable doesn't exists, returns the
// default value
func GetString(key string, defaultValue string) string {
	variable := os.Getenv(key)

	if len(variable) == 0 {
		return defaultValue
	}

	return variable
}
