package test

import "fmt"

type Base struct {
	Colour string
}

type Rabbit struct {
	Base
	ID   int
	Name string
	Age  int
}

func (item *Rabbit) Run() string {
	fmt.Println(item.ID," rabbit is running!!!")
	return " rabbit is running!"
}

func (item *Rabbit) Eat() string {
	fmt.Println(item.ID," rabbit is eating!!!")
	return " rabbit is eating!"
}

func ListArticles() {
	list := make([]string, 3)
	list[0] = "111"
	list[1] = "222"
	list[2] = "333"
	fmt.Println(list)

	//调用同一个包中的其他方法
	Say()
}

func GetComments() {
	commentList := make(map[string]string)
	commentList["edward"] = "nice"
	commentList["star"] = "bad"
	fmt.Println(commentList)
}

func GetMyChan() {
	myChan := make(chan int, 3)
	defer close(myChan)

	myChan <- 3
	fmt.Println(myChan)
}
