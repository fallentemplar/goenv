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
