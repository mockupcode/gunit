package assert

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type T interface {
	Errorf(format string, args ...interface{})
}

func Equal(t T, expected, actual interface{}) bool {
	if !ObjectEqual(expected, actual) {
		return Fail(t, fmt.Sprintf("Expected %s but found %s.", expected, actual))
	}
	return true
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

func Fail(t T, message string) bool {
	t.Errorf("%s %s", "\r"+strings.Join(callerStackTrace(), "\n\r"), message)
	return false
}

func callerStackTrace() (callers []string) {
	for i := 0; ; i++ {
		pc, file, line, ok := runtime.Caller(i)

		if !ok {
			break
		}

		f := runtime.FuncForPC(pc)
		if f == nil {
			break
		}

		path := strings.Split(file, "/")
		file = path[len(path)-1]
		if len(path) > 1 {
			dir := path[len(path)-2]
			if dir != "assert" {
				callers = append(callers, fmt.Sprintf("%s:%d", file, line))
			}
		}

		name := strings.Split(f.Name(), ".")
		function := name[len(name)-1]
		if strings.HasPrefix(function, "Test") {
			break
		}
	}
	return
}
