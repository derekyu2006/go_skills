package go_ticker

import (
	"fmt"
	"time"

	"github.com/derekyu2006/dglog"
)

func TickExec() {
	ticker := time.NewTicker(1 * time.Second)

	i := 0
	for {
		<-ticker.C
		dglog.Infof("abc123344")
		time.Sleep(5 * time.Second)

		i++
		fmt.Println("i = ", i)
		if i == 10 {
			ticker.Stop()
			break
		}
	}
}
