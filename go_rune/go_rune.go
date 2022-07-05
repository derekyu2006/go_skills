package go_rune

import (
	"strconv"

	"github.com/derekyu2006/dglog"
)

func RuneSlice() {
	rowkey := "0000000001293551"
	rowKeyBytes := []rune(rowkey)

	dglog.Infof("rowkeyBytes: %+v", rowKeyBytes)

	dglog.Infof("rune[3:11]: %s", string(rowKeyBytes[3:11]))

	timestamp, err := strconv.ParseInt(string(rowKeyBytes[3:11]), 10, 32) //取[3-10]的timestamp
	if err != nil {
		dglog.Errorf("rowkey parser fail with err: %+v", err)
		return
	}

	dglog.Infof("rowkey: %s, timestamp: %d", rowkey, timestamp)
}
