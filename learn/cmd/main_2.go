package main

func addOne(v *int, initChanel chan int) {
	oldValue := *v
	*v = oldValue + 1
	initChanel <- oldValue + 1
	// *v++
}

// func main() {
// 	value := 0
// 	for range 100 {
// 		value++
// 		go addOne(&value)
// 	}

// 	fmt.Print(value)
// }

// synchroneization
// semaphore
// race condition

// func main() {
// 	value := 0
// 	initChanel := make(chan int, 10)
// 	// javarx
// 	// observer pattern

// 	wg := new(sync.WaitGroup)
// 	mutex := new(sync.Mutex)

// 	for range 100 {
// 		wg.Add(1)

// 		go func() {
// 			defer wg.Done()

// 			mutex.Lock()
// 			defer mutex.Unlock()
// 			addOne(&value, initChanel)

// 			if value == 100 {
// 				close(initChanel)
// 			}
// 		}()
// 	}

// 	// wg.Wait()

// 	// fmt.Print(value)
// 	for {
// 		outputValue, ok := <-initChanel
// 		if !ok {
// 			break
// 		}
// 		fmt.Println(outputValue)
// 	}
// }
