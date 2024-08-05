package FirstGoroutine

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, World!")
}

func main() {
	go sayHello() // This will run as a goroutine
	time.Sleep(1 * time.Second)
	fmt.Println("Goodbye!")
}
