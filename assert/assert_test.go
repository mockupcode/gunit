package assert

import "testing"

func TestEqual(t *testing.T) {
	mockT := new(testing.T)

	if !Equal(mockT, "String", "String") {
		t.Error("String equal must return true")
	}
	if Equal(mockT, "String", "Not String") {
		t.Error("String not equal must return false")
	}
}

func TestObjectEqual(t *testing.T) {
	if !ObjectEqual("String", "String") {
		t.Error("String object equal must return true")
	}
	if ObjectEqual("String", "Not String") {
		t.Error("String object not equal must return false")
	}
}

func TestFail(t *testing.T) {
	mockT := new(testing.T)
	if Fail(mockT, "Message") {
		t.Error("Fail must return false")
	}
}
