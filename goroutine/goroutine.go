package goroutine

import (
	"fmt"
	"time"
)

var chanInt chan int = make(chan int, 10)
var timeout chan bool = make(chan bool)

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
	for{
		select {
		case num :=  <- chanInt:
			fmt.Println("num : ", num)
		case <- timeout:
			fmt.Println("timeout...")
		}

	}
}