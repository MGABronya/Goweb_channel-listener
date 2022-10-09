// @Title  channel-listener
// @Description  用select关键字创建多通道监听器
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-10-10 00:39
package main

import (
	"fmt"
	"time"
)

// @title    foo
// @description   初始化一个通道，并用一个goroutine传入一个值
// @auth      MGAronya（张健）             2022-10-10 00:39
// @param     i int				传入值
// @return    chan int			被传入的通道
func foo(i int) chan int {
	ch := make(chan int)
	// TODO 启用一个goroutine向通道传入值
	go func() { ch <- i }()
	return ch
}

// @title    main
// @description   监视各个通道的数据输出
// @auth      MGAronya（张健）             2022-10-10 00:39
// @param     i int				传入值
// @return    chan int			被传入的通道
func main() {
	ch1, ch2, ch3 := foo(3), foo(6), foo(9)
	ch := make(chan int)

	// TODO 开启一个goroutine监视各个通道数据输出，并收集数据到通道ch里
	go func() {
		// TODO 设置一个计时器，时间到了则发送一个超时信号
		timeout := time.After(1 * time.Second)
		for isTimeout := false; !isTimeout; {
			// TODO 监视通道ch1、ch2、ch3、timeout中的数据输出
			select {
			case v1 := <-ch1:
				ch <- v1
			case v2 := <-ch2:
				ch <- v2
			case v3 := <-ch3:
				ch <- v3
			case <-timeout:
				// TODO 超时
				isTimeout = true
			}
		}
	}()

	// TODO 阻塞主线，取出通道ch中的数据
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}
