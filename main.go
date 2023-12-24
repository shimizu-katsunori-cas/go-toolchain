package main

import (
	"fmt"

	"github.com/shimizu-katsunori-cas/go-toolchain/l1"
	"github.com/shimizu-katsunori-cas/go-toolchain/l2"
	"github.com/shimizu-katsunori-cas/go-toolchain/l3"
)

func main() {
	fmt.Println(l1.Sum(1, 2))
	fmt.Println(l2.Sum(1, 2))
	fmt.Println(l3.Sum(1, 2))
}
