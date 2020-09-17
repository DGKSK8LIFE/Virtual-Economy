package mutate

import (
	"fmt"
	"time"
)

// Pool is the amount of money that is available globally
var Pool int64

// Setup starts the economy
func Setup() {
	Pool = 0
	mutate()
}

func mutate() {
	for {
		// Artificial inflation; will probably make it temporary
		Pool++
		fmt.Println(Pool)
		time.Sleep(1000)
	}
}
