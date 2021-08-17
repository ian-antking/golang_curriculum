package main

import "fmt"
import "flag"

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
	pizzasPtr := flag.Int("pizzas", 1, "number of pizzas to share")
	peoplePtr := flag.Int("people", 1, "number of people")

	flag.Parse()

	fmt.Printf("%+v", sharePizza(*pizzasPtr, *peoplePtr))
}
