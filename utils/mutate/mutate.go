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
	for i := 1; true; i++ {
		// Artificial inflation; will probably make it temporary
		err := db.CreateToken()
		if err != nil {
			panic(err)
		}
		result, err := db.AllTokens()
		if err != nil {
			panic(err)
		}
		count, err := db.TokenCount()
		if err != nil {
			panic(err)
		}
		log.Printf("	Token: %s\n", result["tokenid"])
		if i%5 == 0 {
			fmt.Printf("Total Tokens in circulation: %d\n", count)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
