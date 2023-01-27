package TConsoleInterface

import (
	"fmt"
	"net/http"
)

func CartItem() {
	fmt.Println("Hi, you are here to perform CRUD operations on 'Cart' table")
	fmt.Printf("1.Add To Cart\n2.Get Value From Cart\n3.Delete Item From Cart\n4.Create Cart Reference\n")
	fmt.Println("Please enter your choice")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(err)
	}
	if choice == 1 {
		AddToCart()
	} else if choice == 2 {
		GetCartValue()
	} else if choice == 3 {
		DeleteItemFromCart()
	} else if choice == 4 {
		GetCartReference()
	}
}

func GetCartReference() {
	fmt.Println("YOUR REFERENCE ID WILL BE CREATED BY NOW! PLEASE NOTE IT")
	_, err := http.Post("http://localhost:8089/cart/createreference", "application/json", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func AddToCart() {
	fmt.Println("Please enter the cart reference")
	var ref string
	_, err := fmt.Scanln(&ref)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the valid product id")
	var product_id string
	_, err = fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the quantity of product")
	var quantity string
	_, err = fmt.Scanln(&quantity)
	if err != nil {
		fmt.Println(err)
	}

	url := "http://localhost:8089/addtocart?ref=" + ref + "&product_id=" + product_id + "&quantity=" + quantity

	_, err = http.Post(url, "application/json", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err = fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		Console()
	} else {
		return
	}

}

func GetCartValue() {
	fmt.Println("Please enter the cart reference")
	var ref string
	_, err := fmt.Scanln(&ref)
	if err != nil {
		fmt.Println(err)
	}

	_, err = http.Get("http://localhost:8089/cart/get?ref=" + ref)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err = fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		Console()
	} else {
		return
	}

}

func DeleteItemFromCart() {
	fmt.Println("Please enter the cart reference")
	var ref string
	_, err := fmt.Scanln(&ref)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the valid product id")
	var product_id string
	_, err = fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err)
	}

	url := fmt.Sprintf("http://localhost:8089/deletefromcart?ref=%v&product_id=%v", ref, product_id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Delete done succesfully")

	fmt.Println("Do you want to continue? (yes or no)")
	var cont string
	_, err = fmt.Scanln(&cont)
	if err != nil {
		fmt.Println(err)
	}
	if cont == "yes" {
		Console()
	} else {
		return
	}

}
