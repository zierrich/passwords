package main

import (
	"flag"
	"fmt"

	"github.com/zierrich/passwords"
)

func main() {
	count := flag.Int("n", 10, "number of passwords to generate")
	flag.Parse()

	pass := passwords.New()

	for i := 0; i < *count; i++ {
		fmt.Println(pass.Generate())
	}
}