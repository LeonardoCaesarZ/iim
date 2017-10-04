package crypto

import (
	"fmt"
)

// Init cryption module
func Init() {
	fmt.Println("[module crypto]")
	fmt.Print("read RSA private key from file... ")
	if err := readPriKeyIntoMemroy(); err != nil {
		fmt.Println("[FAIL]")
		panic(err)
	}
	fmt.Println("[OK]")
}
