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
	if !Equal(mockT, nil, nil) {
		t.Error("Nil on expected and actual must return true")
	}
	if Equal(mockT, nil, "String") {
		t.Error("Compair with nil must return false")
	}
}

func TestObjectEqual(t *testing.T) {
	if !ObjectEqual("String", "String") {
		t.Error("String object equal must return true")
	}
	if ObjectEqual("String", "Not String") {
		t.Error("String object not equal must return false")
	}
	if !ObjectEqual(nil, nil) {
		t.Error("Nil on expected and actual must return true")
	}
	if ObjectEqual(nil, "String") {
		t.Error("Compair with nil must return false")
	}
	if !ObjectEqual([]byte("String"), []byte("String")) {
		t.Error("String byte equal must return true")
	}
	if ObjectEqual([]byte("String"), []byte("Not String")) {
		t.Error("String byte not equal must return false")
	}
	if !ObjectEqual(10, 10) {
		t.Error("Number object equal must return true")
	}
	if ObjectEqual(10, 20) {
		t.Error("Number object not equal must return false")
	}
}

func TestFail(t *testing.T) {
	mockT := new(testing.T)
	if Fail(mockT, "Message") {
		t.Error("Fail must return false")
	}
}
