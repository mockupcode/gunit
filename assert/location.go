package assert

import (
	"runtime"
	"strings"
)

func assertLocation(skip int) *location {
	testName, path, line := getTestCallerInfo(skip + 1)
	fileName := getFileName(path)
	return &location{
		Test:     testName,
		FileName: fileName,
		Line:     line,
	}
}

func getTestCallerInfo(skip int) (testName string, path string, line int) {
	testName, path, line = callerInfo(skip + 1)
	for i := skip + 2; !strings.HasPrefix(testName, "Test"); i++ {
		testName, _, _ = callerInfo(i)
	}
	return
}

func callerInfo(skip int) (funcName string, path string, line int) {
	pc, path, line, ok := runtime.Caller(skip + 1)
	if !ok {
		panic("caller skip not pass")
	}
	funcName = getFuncName(pc)
	return
}

func getFuncName(pc uintptr) string {
	fullName := runtime.FuncForPC(pc).Name()
	return substringAfterLast(fullName, ".")
}

func getFileName(path string) string {
	return substringAfterLast(path, "/")
}

func substringAfterLast(str string, separator string) string {
	index := strings.LastIndex(str, separator)
	return str[index+1:]
}
