package main

import "fmt"

func main(){

	x := soma(1, 2, 3)
	z := subtrai(2, 3)
	y := multiplica(10, 10)
	q := divide(20)
	fmt.Println(x, z, y, q)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}

	return total
}

func subtrai(i ...int) int {
	total := 0
	for _, v := range i {
		total = v - total
	}
	return total
}

func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}

func divide(i ...int) int {
	total := 1
	for _, v := range i {
		total = total / v
	}
	return total
}







