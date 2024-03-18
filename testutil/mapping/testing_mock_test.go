package mapping

import (
	"fmt"
	"strings"
)

type testingTMock struct {
	error string
}

func (t *testingTMock) Errorf(format string, args ...interface{}) {
	// the last line is the error message
	lines := strings.Split(strings.TrimSpace(fmt.Sprintf(format, args...)), "\t")

	t.error = lines[len(lines)-1]
}

func (t *testingTMock) FailNow() {
	panic(t.error)
}

func (t *testingTMock) Helper() {
	// do nothing
}
