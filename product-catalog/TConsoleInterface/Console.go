package TConsoleInterface

import (
	"fmt"
)

func Console() {
	fmt.Println("WELCOME TO CONSOLE INTERFACE")
	fmt.Println("Please choose the table of your choice to perform the task")
	fmt.Printf("1.Product_master\n2.Category_master\n3.Inventory\n4.CartItem\n")
	fmt.Println("Enter your choice of table")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(err)
	}

	if choice == 1 {
		Product_master()
	} else if choice == 2 {
		Category_master()
	} else if choice == 3 {
		Inventory()
	} else if choice == 4 {
		CartItem()
	}

}
