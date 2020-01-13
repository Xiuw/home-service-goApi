package main

import (
	"fmt"
	"github.com/Xiuw/home-service-goApi"
	"github.com/thoas/go-funk"
)

func main() {
	chocolates := []string{"perk", "munch", "fivestar"}
	fmt.Println(funk.Contains(chocolates, "perk"))
	fmt.Println("I've ddddddlddhjddjkdd")
	Foo(35)
	Woo(45)
}

func Foo(bar int) int {
	if bar > 0 {
		return 123
	} else {
		return 456
	}
}
