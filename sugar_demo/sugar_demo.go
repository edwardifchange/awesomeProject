package sugar_demo

import "fmt"

func Sugar1(values...string) {
	for _, v := range values {
		fmt.Println("-----> v : ", v)
	}
}

func Sugar2() {
	value := [...]string{"apple", "orange"}
	fmt.Println("-----> value : ", value)
}
