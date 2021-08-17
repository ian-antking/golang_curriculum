package main

import "fmt"

func hello(names ...string) string {
	var name string;
	
	if len(names) > 0 {
		name = names[0]
	} else {
		name = "World"
	}

	return fmt.Sprintf("Hello %v!", name)
}

func main() {
	fmt.Println(hello())
}
