package main

import "fmt"

type pizzaReport struct {
	slicesPerPerson int
	leftoverSlices int
}

func sharePizza(pizzas int, people int) pizzaReport {
	const slicesPerPizza = 8
	var totalSlices int = 8 * pizzas;

	return pizzaReport{ slicesPerPerson: totalSlices/people, leftoverSlices: totalSlices%people}
}

func main() {
	fmt.Println(sharePizza(1, 2))
}
