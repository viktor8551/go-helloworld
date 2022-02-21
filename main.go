package main

import (
	"fmt"
)

func main() {
	bytes := []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}
	pinguin := "\xF0\x9F\x90\xA7"

	fmt.Println(string(bytes), pinguin)
}
