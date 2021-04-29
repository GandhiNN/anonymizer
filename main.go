package main

import (
	"fmt"

	"github.com/GandhiNN/anonymizer/hasher"
)

const STRINGTOHASH = "MyBankPassword"

// main method
func main() {

	// invoke hasher
	hashed := hasher.MD5Hash(STRINGTOHASH)
	fmt.Println(hashed)
}
