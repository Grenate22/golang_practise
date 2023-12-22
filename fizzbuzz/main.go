package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Printf("%d is fizzbuzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d is Fizz\n", i)

		} else if i%5 == 0 {
			fmt.Printf("%d is buzz\n", i)
		} else {
			fmt.Println(i)
		}
	}
}
