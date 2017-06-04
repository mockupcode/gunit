package assert

import (
	"bytes"
	"fmt"
	"reflect"
)

type T interface {
	Fail()
}

type logFacade struct {
	t      T
	logger logger
}

type assertion struct {
	logFacade *logFacade
}

func GetAssertion(t T) *assertion {
	return &assertion{&logFacade{t, theLogger}}
}

func fail(log *logFacade, message string) bool {
	log.logger.Log(assertLocation(2), message)
	log.t.Fail()
	return false
}

func objectEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}
	if exp, ok := expected.([]byte); ok {
		if act, ok := actual.([]byte); ok {
			return bytes.Equal(exp, act)
		}
	}
	return reflect.DeepEqual(expected, actual)
}

func (a *assertion) Equal(expected, actual interface{}) bool {
	if !objectEqual(expected, actual) {
		return fail(a.logFacade, fmt.Sprintf("Expected %v but found %v.", expected, actual))
	}
	return true
}
