package assert

import (
	"fmt"
	"io"
	"os"
)

type location struct {
	Test     string
	FileName string
	Line     int
}

type logger interface {
	Log(location *location, message string)
}

type loggerImpl struct {
	writer       io.Writer
	prevTestName string
	prevTestLine int
}

var theLogger logger = &loggerImpl{writer: os.Stdout}

const (
	header = "--- FAIL: %s\n\t%s:%d: %s\n"
	body   = "\t%s:%d: %s\n"
)

func (logger *loggerImpl) Log(location *location, message string) {
	if logger.prevTestName != location.Test {
		fmt.Fprintf(logger.writer, header, location.Test, location.FileName, location.Line, message)
	} else if logger.prevTestLine != location.Line {
		fmt.Fprintf(logger.writer, body, location.FileName, location.Line, message)
	}
	logger.prevTestName = location.Test
	logger.prevTestLine = location.Line
}
