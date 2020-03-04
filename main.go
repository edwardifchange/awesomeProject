package main

import (
	"awesomeProject/interface_demo"
	"awesomeProject/point_demo"
	"fmt"
	"reflect"
)

func main() {
	//Test for make()
	//makeSlice()
	//makeMap()
	//makeChan()

	//Test for new()
	//NewMap()

	//Test for append
	//appendElementForSlice()

	//Test for copy
	//copyForClice()

	//Test for delete
	//deleteFromMap()

	//Test for panic() and recover()
	//receivePanic()

	//Test for len() and cap()
	//getLen()

	//Test for close()
	//closeChan()

	//createArray()

	//struct_demo.TestForStruct()
	//test.ListArticles()

	//Test for Run()
	//dog := new(struct_demo.Dog)
	//dog.ID = 1
	//dog.Name = "DuDu"
	//dog.Age = 4
	//dog.Colour = "grey"
	//dog.Run()
	//dog.Eat()

	//测试接口定义变量
	//var test interface_demo.Behavior
	//test = new(struct_demo.Cat)
	//test.Run()

	//dog := new(struct_demo.Dog)
	//dog.ID = 1
	//dog.Name = "BeiBei"
	//dog.Age = 15
	//dog.Colour = "yellow"
	//action(dog)
	//
	//cat := new(struct_demo.Cat)
	//cat.ID = 2
	//cat.Name = "DuDu"
	//cat.Age = 1
	//cat.Colour = "grey"
	//action(cat)

	//action(new(struct_demo.Cat))

	//并发
	//fmt.Printf("cpu num = %d", runtime.NumCPU())
	//runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	//go goroutine.Loop()
	//go goroutine.Loop()
	//time.Sleep(time.Second * 5)

	////启动发送数据的协程
	//go goroutine.Send()
	////启动接收数据的协程
	//go goroutine.Receive()
	//time.Sleep(time.Second * 10)

	////协程同步
	//goroutine.Read()
	//go goroutine.Write()
	//goroutine.WG.Wait()
	//fmt.Println("all done !")
	//time.Sleep(time.Second * 60)

	//指针
	//point_demo.TestPoint()

	point_demo.TestPointArr()
}

func action(test interface_demo.Behavior) string {
	test.Run()
	test.Eat()
	return "suc"
}

func createArray() {
	//定义赋值一个数组方式一
	//var a [10]int
	//a[0] = 1
	//fmt.Printf("%d\n", a)

	// 定义赋值一个数组方式二
	//var b [10]int = [10]int{1, 2, 3}
	//fmt.Printf("%d\n", b)

	// 定义赋值一个数组方式三
	c := [10]int{1, 2, 3}
	//c := [...]int{}
	fmt.Printf("%d\n", c)
	//
	//test := []string{"dog"}
	//fmt.Println(test)
}

//makeSlice 创建切片
func makeSlice() {
	mSlice := make([]string, 3) //长度是3
	mSlice[0] = "dog"
	mSlice[1] = "cat"
	mSlice[2] = "tiger"

	fmt.Println(mSlice)
}

//makeMap 创建map
func makeMap() {
	mMap := make(map[int]string)
	mMap[10] = "dog"
	mMap[100] = "cat"

	fmt.Println(mMap)
}

//makeChan 创建chan
func makeChan() {
	mChan := make(chan int, 3) //管道中的 缓存，容量 是3，如不写，则为没有缓存的chan
	close(mChan)
}

func NewMap() {
	//mMap := make(map[int]string)
	mMap := new(map[int]string)
	fmt.Println(reflect.TypeOf(mMap))
}

//向切片中添加元素并打印长度，容量
func appendElementForSlice() {
	mIDslice := make([]string, 2)
	mIDslice[0] = "id-1"
	mIDslice[1] = "id-2"
	fmt.Println("len=", len(mIDslice))
	fmt.Println("cap=", cap(mIDslice))

	mIDslice = append(mIDslice, "id-3")
	fmt.Println(mIDslice)
	fmt.Println("After len=", len(mIDslice))
	fmt.Println("After cap=", cap(mIDslice))
}

//复制切片
func copyForClice() {
	mIDsliceDst := make([]string, 2)
	mIDsliceDst[0] = "id-dst-1"
	mIDsliceDst[1] = "id-dst-2"

	mIDsliceSrc := make([]string, 3)
	mIDsliceSrc[0] = "id-src-1"
	mIDsliceSrc[1] = "id-src-2"
	mIDsliceSrc[2] = "id-src-3"

	copy(mIDsliceDst, mIDsliceSrc) //copy不会给目标切片扩容
	fmt.Println(mIDsliceDst)
}

//删除map中的数据
func deleteFromMap() {
	mIDMap := make(map[int]string)
	mIDMap[0] = "id-1"
	mIDMap[1] = "id-2"
	delete(mIDMap, 0)
	fmt.Println(mIDMap)
}

//异常捕获
func receivePanic() {
	defer coverPanic()
	//panic("I am Panic")
	//panic(errors.New("I am Panic"))
	panic(1)
}

func coverPanic() {
	message := recover()
	switch message.(type) {
	case string:
		fmt.Println("string message: ", message)
	case error:
		fmt.Println("error message: ", message)
	default:
		fmt.Println("unknown message: ", message)
	}
}

//获取长度和容量
func getLen() {
	mIDsliceDst := make([]string, 2, 5)
	mIDsliceDst[0] = "id-dst-1"
	mIDsliceDst[1] = "id-dst-2"
	mIDsliceDst = append(mIDsliceDst, "id-dst-3")
	mIDsliceDst = append(mIDsliceDst, "id-dst-4")
	mIDsliceDst = append(mIDsliceDst, "id-dst-5")
	mIDsliceDst = append(mIDsliceDst, "id-dst-6")

	fmt.Println(len(mIDsliceDst))
	fmt.Println(cap(mIDsliceDst))
}

//创建和关闭chan
func closeChan() {
	mChan := make(chan int, 1)
	mChan <- 1
	defer close(mChan) //防止后面忘记关闭，所以加上defer，保证该方法最后有执行
}
