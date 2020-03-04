package point_demo

import "fmt"

func TestPoint() {
	var count int = 30
	var countPoint *int
	var countPoint1 *int
	countPoint = &count
	//countPoint++ go语言不支持指针运算

	fmt.Printf("count 变量的地址： %x \n", &count)
	if countPoint != nil {
		fmt.Printf("countPoint 变量存储的地址： %x \n", countPoint)

	}
	fmt.Printf("countPoint 指针指向地址的值： %d \n", *countPoint)
	fmt.Printf("countPoint1 变量存储的地址： %d \n", countPoint1)
	fmt.Println("countPoint1 指针指向地址的值：", countPoint1)
}

func TestPointArr() {
	//指针数组
	a, b := 1, 2
	pointArr := [...]*int{&a, &b}
	fmt.Println("指针数组 : ", pointArr)

	//数组指针
	arr := [...]int{3, 4, 5}
	arrPoint := &arr
	fmt.Println("数组指针 : ", arrPoint)
}
