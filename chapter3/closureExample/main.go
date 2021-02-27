package main

import "fmt"

func main() {

	gene := generator()

	for i := 0; i < 5; i++ {
		fmt.Println(gene(), "\t")
	}

}

func generator() func() int {
	var i = 0
	return func() int {
		i++
		return i
	}
}
