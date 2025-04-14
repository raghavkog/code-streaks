package codedaily

import (
	"fmt"
	"sync"
)

func runTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d done\n", id)
}

func Main3() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go runTask(i, &wg)
	}

	wg.Wait()
	fmt.Println("All tasks completed")
}
