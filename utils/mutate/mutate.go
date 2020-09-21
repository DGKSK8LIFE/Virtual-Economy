package mutate

import (
	"economy/utils/db"
	"fmt"
	"log"
	"time"
)

// Mutate mutates the economy
func Mutate() {
	fmt.Println("Creating tokens:")
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
		log.Printf("	Token: %s\n", result["tokenid"])
		time.Sleep(1000 * time.Millisecond)
	}
}
