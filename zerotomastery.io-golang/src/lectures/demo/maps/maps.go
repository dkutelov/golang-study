package main

import "fmt"

func main() {
	// Maps are unordered!
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 12
	shoppingList["bread"] += 1 // do not need to initialise since default is 0
	shoppingList["eggs"] += 1

	fmt.Println(shoppingList)

	//delete
	delete(shoppingList, "milk")
	fmt.Println(shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	// cereal is default value zero if found is false

	if !found {
		fmt.Println("nope")
	} else {
		fmt.Println("yep", cereal)
	}

	// Number of existing keys
	productCount := len(shoppingList)
	totalItems := 0
	// range - order will be different each time
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("Number of products", productCount)
	fmt.Println("Quantity of products", totalItems)

}
