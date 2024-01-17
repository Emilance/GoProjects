package main

import "fmt"

var a = 2

func increment() {
	a += 1
}
func decrement() {
	a -= 1
}

func main() {
	go increment()
	go decrement()

	fmt.Println("New value of a :", a)

}

// The code above can lead to unpredictable result
// The race condition here arises due to concurrent access to the shared variable a without
// proper synchronization.
// The increment and decrement functions are executed concurrently as goroutines, and both
// may access and modify the variable a simultaneously.
// The order of execution is not guaranteed in concurrent programs, so it's not
// predictable whether the increment or decrement operation will happen first. Consequently,
//  the final output value of a is unpredictable.
