package goroutine

import (
	"fmt"
	"time"
)

func Loop() {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Microsecond * 10)
		fmt.Printf("%d, ", i)
	}
}

