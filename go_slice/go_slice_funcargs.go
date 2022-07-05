package go_slice

import "github.com/derekyu2006/dglog"

func FuncSliceArgs(s []string) {
	dglog.Debugf("input s: %+v", s)
	s = append(s, "123456")
}

func MainFunc() {
	s := []string{"hello world"}
	FuncSliceArgs(s)
	dglog.Debugf("after s: %+v", s)
}
