package main

import (
	"log"
	"os"
)

func main() {
	fp, err := os.Open("crazy.csv")
	if err != nil {
		log.Fatal(err)
	}
}
