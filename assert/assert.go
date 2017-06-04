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

type Assertion struct {
	logFacade *logFacade
}

func GetAssertion(t T) *Assertion {
	return &Assertion{&logFacade{t, theLogger}}
}

func Fail(log *logFacade, message string) bool {
	log.logger.Log(assertLocation(2), message)
	log.t.Fail()
	return false
}

func ObjectEqual(expected, actual interface{}) bool {
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

func (a *Assertion) Equal(expected, actual interface{}) bool {
	if !ObjectEqual(expected, actual) {
		return Fail(a.logFacade, fmt.Sprintf("Expected %v but found %v.", expected, actual))
	}
	return true
}
