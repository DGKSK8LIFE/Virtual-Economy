package main

import (
	"economy/utils/db"
	"economy/utils/mutate"
)

func main() {
	db.Open()
	mutate.Setup()
}
