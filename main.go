package main

import (
	"fmt"

	"github.com/minio/sha256-simd"
)

func main() {
	hash := sha256.New()
	fmt.Printf("%v\n", hash)
}
