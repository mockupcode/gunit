package assert

import (
	"bytes"
	"testing"
)

func TestEqual(t *testing.T) {
	mockT := new(testing.T)
	buffer := &bytes.Buffer{}
	assert := &assertion{logFacade{mockT, &loggerImpl{writer: buffer}}}

	if !assert.Equal("String", "String") {
		t.Error("Equal must return true")
	}
	if assert.Equal("String", "Not String") {
		t.Error("Equal must return true")
	}
	if !assert.Equal(nil, nil) {
		t.Error("Equal must return true")
	}
	if assert.Equal(nil, "String") {
		t.Error("Equal must return true")
	}
	if !assert.Equal([]byte("String"), []byte("String")) {
		t.Error("Equal must return true")
	}
	if assert.Equal([]byte("String"), []byte("Not String")) {
		t.Error("Equal must return true")
	}
	if !assert.Equal(10, 10) {
		t.Error("Equal must return true")
	}
	if assert.Equal(10, 20) {
		t.Error("Equal must return true")
	}
	if !assert.Equal(10.5, 10.5) {
		t.Error("Equal must return true")
	}
	if assert.Equal(10.5, 20.5) {
		t.Error("Equal must return true")
	}
}

func TestGetAssertion(t *testing.T) {
	mockT := new(testing.T)
	assert := GetAssertion(mockT)
	if (assertion{}) == assert {
		t.Error("GetAssertion must not zero value")
	}
}

func TestFailIsPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	mockT := new(testing.T)
	buffer := &bytes.Buffer{}
	fail(logFacade{mockT, &loggerImpl{writer: buffer}}, "message")
}
