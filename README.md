# GUnit [![GoDoc](https://godoc.org/github.com/mockupcode/gunit/assert?status.svg)](https://godoc.org/github.com/mockupcode/gunit/assert) [![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/mockupcode/gunit/blob/master/LICENSE) [![Build Status](https://travis-ci.org/mockupcode/gunit.svg?branch=master)](https://travis-ci.org/mockupcode/gunit) [![Coverage Status](https://coveralls.io/repos/github/mockupcode/gunit/badge.svg)](https://coveralls.io/github/mockupcode/gunit)

Assertion library for golang!

Getting started
---------------
To get the package, execute:

	go get github.com/mockupcode/gunit 

To import this package, add the following line to your code:

	import "github.com/mockupcode/gunit/assert"

And just start using it.
```go
import (
	"testing"

	"github.com/mockupcode/gunit/assert"
)

func TestAssertion(t *testing.T){
	assert := assert.GetAssertion(t)
	assert.Equal("string","string")
}
```
License
-------
GUnit is a free software and may be used under the terms specified in the [LICENSE](LICENSE) file.
