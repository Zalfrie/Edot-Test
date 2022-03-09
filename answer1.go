package main

import "fmt"

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		s := ""
		if i % 3 == 0 {
			s += "Fizz"
		}
		if i % 5 == 0 {
			s += "Buzz"
		}
		if s != "" {
			fmt.Println(s)
		} else {
			fmt.Println(i)
		}
	}
}