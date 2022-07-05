package go_string

import "github.com/derekyu2006/dglog"

func FuncStringArgs(s string) {
	dglog.Debugf("input s: %+v", s)
	s = "123456"
}

func MainFunc() {
	s := "hello world"
	FuncStringArgs(s)
	dglog.Debugf("after s: %+v", s)
}
