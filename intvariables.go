package goenv

import (
	"os"
	"strconv"
)

//GetInt recovers an environment variable and casts it to a int
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
