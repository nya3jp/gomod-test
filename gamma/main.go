package main

import (
	"fmt"

	"github.com/nya3jp/gomod-test/alpha"
	"github.com/nya3jp/gomod-test/beta"
)

func main() {
	fmt.Printf("alpha.Message = %s\n", alpha.Message)
	fmt.Printf("beta.Message = %s\n", beta.Message)
}
