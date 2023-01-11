package L1

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	numbers := []uint8{2, 4, 6, 8, 10}
	go func() {
		for i := 0; i < len(numbers); i++ {
			fmt.Println(i)
		}
	}()

	elapsedTime := time.Since(start)

	fmt.Println("Total Time For Execution: " + elapsedTime.String())

	time.Sleep(time.Second)
}
