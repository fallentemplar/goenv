package test

import (
	"os"
	"testing"

	"github.com/fallentemplar/goenv"
)

func TestGetInt(t *testing.T) {
	os.Setenv("banana", "")
	result := goenv.GetInt("banana", 0)

	if result != 0 {
		t.Errorf("Unexpected result")
	}

	os.Setenv("banana", "1547")

	result = goenv.GetInt("banana", 1547)

	if result != 1547 {
		t.Errorf("Unexpected result")
	}
}

func TestGetBool(t *testing.T) {
	//Getting a non-existent variable
	os.Setenv("isDevelopment", "")
	expectedResult := true

	result := goenv.GetBool("isDevelopment", true)

	if result != expectedResult {
		t.Errorf("Unexpected result. Expected: %t\tReceived: %t", expectedResult, result)
	}

	//Getting an existent variable
	os.Setenv("isDevelopment", "0")
	expectedResult = false
	result = goenv.GetBool("isDevelopment", false)

	if result != expectedResult {
		t.Errorf("Unexpected result. Expected: %t\tReceived: %t", expectedResult, result)
	}

	//Getting a non-existent variable with different defaultValue
	os.Setenv("isDevelopment", "1")
	expectedResult = true
	result = goenv.GetBool("isDevelopment", false)

	if result != expectedResult {
		t.Errorf("Unexpected result. Expected: %t\tReceived: %t", expectedResult, result)
	}
}

func TestGetString(t *testing.T) {
	//Getting a non-existent variable
	os.Setenv("environment_name", "")

	expectedResult := "development"
	result := goenv.GetString("environment_name", "development")

	if result != expectedResult {
		t.Errorf("Unexpected result: Expecting: %s\tReceived: %s", expectedResult, result)
	}

	//Getting an existent variable
	os.Setenv("environment_name", "development")

	expectedResult = "development"
	result = goenv.GetString("environment_name", "development")

	if result != expectedResult {
		t.Errorf("Unexpected result: Expecting: %s\tReceived: %s", expectedResult, result)
	}

	//Getting an existent variable with another default value
	os.Setenv("environment_name", "development")

	expectedResult = "development"
	result = goenv.GetString("environment_name", "production")

	if result != expectedResult {
		t.Errorf("Unexpected result: Expecting: %s\tReceived: %s", expectedResult, result)
	}
}

func TestGithubActions(t *testing.T) {
	// Getting a existent variable but expecting another value
	// This is a temp test for verifying the GH action is runing right
	os.Setenv("environment_name", "github")

	expectedResult := "development"
	result := goenv.GetString("environment_name", "development")

	if result != expectedResult {
		t.Errorf("Unexpected result: Expecting: %s\tReceived: %s", expectedResult, result)
	}
}
