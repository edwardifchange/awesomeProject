package goroutine

import (
	"fmt"
	"sync"
	"time"
)

var chanInt chan int = make(chan int, 10)
var timeout chan bool = make(chan bool)
var WG sync.WaitGroup

func Loop() {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Microsecond * 10)
		fmt.Printf("%d, ", i)
	}
}

//发送数据到chan
func Send() {
	//time.Sleep(time.Second * 1)
	chanInt <- 1
	time.Sleep(time.Second * 1)
	chanInt <- 2
	time.Sleep(time.Second * 1)
	chanInt <- 3
	time.Sleep(time.Second * 2)
	timeout <- true
}

//接受chan中的数据
func Receive() {
	//num := <- chanInt
	//fmt.Println("num : ", num)
	//
	//num = <- chanInt
	//fmt.Println("num : ", num)
	//
	//num = <- chanInt
	//fmt.Println("num : ", num)

	//应用场景：多个协程从文件中读取，一个协程接收读取到的数据
	for {
		select {
		case num := <-chanInt:
			fmt.Println("num : ", num)
		case <-timeout:
			fmt.Println("timeout...")
		}

	}
}

//协程同步 场景：10个协程读取文件，1个协程写入文件，要等写入文件完成之后，在进行后续的操作

//Add(delta int) 添加携程记录
//Done() 移除协程记录
//Wait() 同步等待所有记录的协程全部结束

//读取数据
func Read() {
	for i := 0; i < 3; i++ {
		WG.Add(1)
	}
}

//写入数据
func Write() {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println("Done ---> ", i)
		WG.Done()
	}
}
