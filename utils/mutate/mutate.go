package mutate

import (
	"economy/utils/db"
	"time"
)

// Mutates the economy
func Mutate() {
	for {
		// Artificial inflation; will probably make it temporary
		err := db.CreateToken()
		if err != nil {
			panic(err)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
