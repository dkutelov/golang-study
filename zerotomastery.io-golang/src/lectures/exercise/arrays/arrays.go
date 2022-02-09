//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name string
	price float64
}


func getListInfo(shoppingList [4]Product) {
	productCount := 0
	lastProduct := Product{}
	var totalCost float64

	for i := 0; i < len(shoppingList); i++ {
		currentProduct := shoppingList[i]

		if currentProduct != (Product{}) {
			productCount += 1
			lastProduct = currentProduct
			totalCost += currentProduct.price
		}
	}

	fmt.Println("The last item on the list is", lastProduct)
	fmt.Println("The total number of items is", productCount)
	fmt.Println("The total cost of the products is", totalCost)
}


func main() {
	shoppingList := [4]Product {
		{name: "item1", price: 1.5},
		{name: "item2", price: 2.5},
		{name: "item3", price: 3.5},
	}

	getListInfo(shoppingList)
	
	shoppingList[3] = Product{name: "item 4", price: 4.5}

	getListInfo(shoppingList)
}
