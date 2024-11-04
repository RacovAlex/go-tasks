package concurrency

import (
	"fmt"
	"sync"
)

// Задача: есть канал int-ов, куда присылаются N значений. Сделать K обработчиков, которые достанут значения,
// сделают +1 и положат в исходящий канал. Потом надо прочитать значения из исходящего канала и напечатать.

func WorkerPool(in <-chan int) {
	const workers = 10
	wg := sync.WaitGroup{}
	out := make(chan int, workers)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var work int
			for v := range in {
				work = v
				out <- work + 1
			}
			//for {
			//	select {
			//	case work = <-in:
			//		out <- work + 1
			//	default:
			//		time.Sleep(time.Millisecond * 10)
			//	}
			//}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for v := range out {
		fmt.Println(v)
	}

}
