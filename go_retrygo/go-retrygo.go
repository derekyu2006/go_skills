package main

import (
	"errors"
	"fmt"
	"time"
)

func retryFunc() error {
	return fmt.Errorf("test retry")
}

//重试，限制次数
func RetryTimesDo(name string, tryTimes int, sleep time.Duration, callback func() error) (err error) {
	for i := 1; i <= tryTimes; i++ {
		err = callback()
		if err == nil {
			return nil
		}
		fmt.Printf("[%v]失败，第%v次重试, 错误信息:%s \n", name, i, err)
		time.Sleep(sleep)
	}

	err = fmt.Errorf("[%v]失败，共重试%d次, 最近一次错误:%s \n", name, tryTimes, err)
	fmt.Println(err)
	return err
}

//重试，限制时间
func RetryDurations(name string, max time.Duration, sleep time.Duration, callback func() error) (err error) {
	t0 := time.Now()
	i := 0
	for {
		err = callback()
		if err == nil {
			return
		}
		delta := time.Now().Sub(t0)
		if delta > max {
			fmt.Printf("[%v]失败，超过最大时间%s, 共重试%d次, 最近一次错误: %s \n", name, max, i, err)
			return err
		}
		time.Sleep(sleep)
		i++
		fmt.Printf("[%v]失败，第%v次重试, 错误信息:%s \n", name, i, err)
	}
}

func main() {
	//示例任务函数，只会失败
	var task = func() error {
		return errors.New("task error")
	}

	fmt.Println("RetryTimes")
	//设定3秒最大重试次数，最大3次，1秒重试一次
	RetryTimesDo("测试任务", 3, time.Second, task)

	fmt.Println("RetryDurations")
	//设定3秒最大重试期限，最多3秒，1秒重试一次
	RetryDurations("测试任务", 3*time.Second, time.Second, task)
}
