package equal

import (
	"testing"
)

func TestLineNumbers(t *testing.T) {
	Equal(t)
	Equal(t, "foo")
	Equal(t, "foo", "foo")
	Equal(t, "foo", "foo", "msg!")
	Equal(t, "foo", "foo", "msg!", "msg")
	Equal(t, "foo", "foo", "msg!", "msg", "haha")
	//Equal(t, "foo", "bar", "this should blow up")
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, "foo", "bar", "msg!")
	//NotEqual(t, "foo", "foo", "this should blow up")
}
