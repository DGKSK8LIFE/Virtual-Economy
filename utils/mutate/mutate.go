package mutate

import (
	"economy/utils/db"
	"fmt"
	"time"
)

// Mutate mutates the economy
func Mutate() {
	for {
		// Artificial inflation; will probably make it temporary
		err := db.CreateToken()
		if err != nil {
			panic(err)
		}
		result, err := db.AllTokens()
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
		time.Sleep(1000 * time.Millisecond)
	}
}
