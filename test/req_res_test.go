package test

import (
	"fmt"
	"testing"
	"time"
)

func Test_Out(t *testing.T) {
	c := make(chan int)
	defer close(c)

	go func() {
		c <- 3 + 4
		fmt.Println("------999999")
	}()
	i := <-c
	fmt.Println(i)
}

func Test_Channel(t *testing.T) {

	/*1.简单的channel队列现进先出特性
	 *在主线程上建立单个协程，完成某项工作后，关闭该管道,关闭与否为可选项。
	 *停止一段阻塞有两种方式，chanInt内能读取到值，即非空;chanInt被关闭
	 */
	if false {
		var chanInt chan int = make(chan int, 10)
		go func() {
			fmt.Println("新线程完成了某项工作")
			defer close(chanInt)
			chanInt <- 1
			chanInt <- 2
			chanInt <- 3
			chanInt <- 4

		}()

		x := <-chanInt
		fmt.Println(x)
		x = <-chanInt
		fmt.Println(x)
		x = <-chanInt
		fmt.Println(x)
		x = <-chanInt
		fmt.Println(x)
	}

	/*2.简单的select使用
	 *使用select只会在符合的语句上走一次，即使两个管道都赋值了，先监测到chan2传值以后，select走完case2就已经结束了
	 *管道对象是引用类型，方法传值并不是值的拷贝，而是可以确切地修改队列内容
	 */
	if true {
		chan1 := make(chan int)
		chan2 := make(chan int)
		go test2(chan1, chan2)
		for {
			select {
			case y := <-chan1:
				fmt.Println("chan1收到值", y)
			case x := <-chan2:
				fmt.Println("chan2收到值", x)
			default:
				{
					fmt.Println("结束了")
					break
				}
			}
		}

	}

	/*如何让主gorouting控制单协程结束
	 */
	if false {
		stopFlag := make(chan bool)
		out := make(chan bool)

		go func() {
			defer close(out)
			fmt.Println("完成任务ing")

			select {
			case <-stopFlag:
				{
					fmt.Println("任务结束")
				}
			}
		}()
		stopFlag <- true
		<-out
	}

	/*
	 *select没有满足条件时会被阻塞,一直没有注值就会死锁
	 *引入超时机制，5秒注值显然大于了超时设定3秒，所以超时了
	 */
	if true {
		chanInt := make(chan int, 1)
		out := make(chan int, 1)
		go func() {
			go func() {
				time.Sleep(5 * time.Second)
				chanInt <- 5

			}()
			select {
			case x := <-chanInt:
				fmt.Println(x)
			case <-time.After(3 * time.Second):
				fmt.Println("超时了")
			}
			out <- 1
		}()

		<-out
	}

}

func test2(a chan int, b chan int) {
	b <- 6
	b <- 7
	a <- 5
}
