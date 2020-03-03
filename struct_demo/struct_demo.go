package struct_demo

import "fmt"

type Animal struct {
	Colour string
}

type Dog struct {
	Animal
	ID   int
	Name string
	Age  int
}

type Cat struct {
	Animal
	ID   int
	Name string
	Age  int
}

//创建结构体的三种方法
func TestForStruct() {
	//方法1
	//var dog Dog
	//dog.ID = 1
	//dog.Name = "DuDu"
	//dog.Age = 3

	//方法2
	//dog := Dog{1, "Dudu", 3}

	//方法3
	dog := new(Dog)
	dog.ID = 1
	dog.Name = "DuDu"
	dog.Age = 4

	fmt.Println(dog)
}

func (item *Dog) Eat() string {
	fmt.Println("delicious!")
	return "delicious!"
}

func (item *Dog) Run() string {
	fmt.Println("ID : ", item.ID, " colour : ", item.Colour, " ,", item.Name, " is running!")
	return "dog is running!"
}

//Cat
func (item *Cat) Eat() string {
	fmt.Println("Cat --> delicious!")
	return "Cat --> delicious!"
}

func (item *Cat) Run() string {
	fmt.Println("ID : ", item.ID, " colour : ", item.Colour, " , ", item.Name, " is running!")
	return "Cat is running!"
}