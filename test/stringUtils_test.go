package test

import (
	"testing"

	StringUtils "github.com/NuttayotSukkum/masterlib/validations/stringUtils"
)

func TestIsBlank(t *testing.T) {
	result := StringUtils.IsBlank(" ")
	if result != true {
		t.Errorf("Expected true, but got %v", result)
	}
}

func TestIsEmpty(t *testing.T) {
	var val *string = nil
	result := StringUtils.IsEmpty(val)
	if result != true {
		t.Errorf("Expected true, but got %v", result)
	}
}

func TestIsNotBlank(t *testing.T) {
	result := StringUtils.IsNotBlank("data")
	if result != false {
		t.Errorf("Expected false, but got %v", result)
	}
}
func TestIsNotEmpty(t *testing.T) {
	var val = ""
	result := StringUtils.IsNotEmpty(&val)
	if result != false {
		t.Errorf("Expected false, but got %v", result)
	}
}
