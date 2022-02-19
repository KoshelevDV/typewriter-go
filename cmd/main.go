package main

import (
	"os"

	"github.com/koshelevdv/typewriter-go/internal/typeWriter"
)

func main() {
	typeWriter.TypeWriter(os.Args[1], os.Args[2], 15)
}
