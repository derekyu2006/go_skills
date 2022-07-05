package go_range

import (
	"fmt"

	"github.com/derekyu2006/dglog"
)

func RangeNil() {
	// no nil
	var datas []string = nil

	for _, ele := range datas {
		fmt.Println(ele)
	}

	dglog.Infof("456 helloword")
}
