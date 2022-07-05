package go_gcache

import (
	"time"

	"github.com/bluele/gcache"
	"github.com/derekyu2006/dglog"
)

func GCacheOne() {
	gc := gcache.New(20).
		LRU().
		Build()
	gc.SetWithExpire("key", "ok", time.Second*2)
	value, _ := gc.Get("key")
	dglog.Infof("Get: %+v", value)

	// Wait for value to expire
	time.Sleep(time.Second * 3)

	value, err := gc.Get("key")
	if err != nil {
		dglog.Infof("Get: %+v", err)
		return
	}
	dglog.Infof("Get: %+v", value)
}
