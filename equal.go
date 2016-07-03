package equal

import (
	"fmt"
	"github.com/kr/pretty"
	"reflect"
	"runtime"
	"testing"
)

func assert(t *testing.T, result bool, f func(), cd int) {
	if !result {
		_, file, line, _ := runtime.Caller(cd + 1)
		t.Errorf("%s:%d", file, line)
		f()
		t.FailNow()
	}
}

func equal(t *testing.T, exp, got interface{}, cd int, args ...interface{}) {
	fn := func() {
		for _, desc := range pretty.Diff(exp, got) {
			t.Error("!", desc)
		}
		if len(args) > 0 {
			t.Error("!", " -", fmt.Sprint(args...))
		}
	}
	result := reflect.DeepEqual(exp, got)
	assert(t, result, fn, cd+1)
}

func tt(t *testing.T, result bool, cd int, args ...interface{}) {
	fn := func() {
		t.Errorf("!  Failure")
		if len(args) > 0 {
			t.Error("!", " -", fmt.Sprint(args...))
		}
	}
	assert(t, result, fn, cd+1)
}

func T(t *testing.T, result bool, args ...interface{}) {
	tt(t, result, 1, args...)
}

func Tf(t *testing.T, result bool, format string, args ...interface{}) {
	tt(t, result, 1, fmt.Sprintf(format, args...))
}

func Equal(t *testing.T, args ...interface{}) {
	if t == nil {
		t = &testing.T{}
	}
	length := len(args)
	if length <= 1 {
		return
	}
	for i := 0; i < length/2; i += 1 {
		if length%2 == 1 {
			equal(t, args[2*i], args[2*i+1], 1, args[length-1])
		} else {
			equal(t, args[2*i], args[2*i+1], 1)
		}
	}
}

func Equalf(t *testing.T, format string, args ...interface{}) {
	length := len(args)
	if length <= 1 {
		return
	}
	for i := 0; i < length/2; i += 1 {
		if length%2 == 1 {
			equal(t, args[2*i], args[2*i+1], 1, fmt.Sprintf(format, args[length-1]))
		} else {
			equal(t, args[2*i], args[2*i+1], 1)
		}
	}
}

func NotEqual(t *testing.T, args ...interface{}) {
	length := len(args)
	if length <= 1 {
		return
	}
	for i := 0; i < length/2; i += 1 {
		fn := func() {
			t.Errorf("!  Unexpected: <%#v>", args[2*i])
			if length%2 == 1 {
				t.Error("!", " -", fmt.Sprint(args[length-1]))
			}
		}
		result := !reflect.DeepEqual(args[2*i], args[2*i+1])
		assert(t, result, fn, 1)
	}

}

func Panic(t *testing.T, err interface{}, fn func()) {
	defer func() {
		equal(t, err, recover(), 3)
	}()
	fn()
}
