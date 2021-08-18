package main

import "fmt"
import "flag"


func sharePizza(pizzas int, people int) (slicesPerPerson, leftoverSlices int) {
	const slicesPerPizza = 8
	var totalSlices int = 8 * pizzas;
	slicesPerPerson = totalSlices/people
	leftoverSlices = totalSlices%people
	return 
}

func main() {
	pizzasPtr := flag.Int("pizzas", 1, "number of pizzas to share")
	peoplePtr := flag.Int("people", 1, "number of people")

	flag.Parse()

	fmt.Println(sharePizza(*pizzasPtr, *peoplePtr))
}
