package TConsoleInterface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Inventory() {
	fmt.Println("Hi Welcome to our inventory section,please feel free to perform CRUD operations on 'Inventory' table")
	fmt.Printf("1.Add\n2.Get\n3.Update\n4.Delete\n")
	fmt.Println("Please enter your choice")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(err)
	}
	if choice == 1 {
		AddInventory()
	} else if choice == 2 {
		GetInventory()
	} else if choice == 3 {
		UpdateInventory()
	} else if choice == 4 {
		DeleteInventory()
	}
}

func AddInventory() {
	fmt.Println("Please enter the valid product id")
	var product_id int
	_, err := fmt.Scanf("%d", &product_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the Quantity for product")
	var quantity int
	_, err = fmt.Scanf("%d", &quantity)
	if err != nil {
		fmt.Println(err)
	}

	own_data := fmt.Sprintf("{\"product_id\":%v,\"quantity\":%v}", product_id, quantity)

	byte_data := []byte(own_data)

	_, err = http.Post("http://localhost:8089/addinventory", "application/json", bytes.NewBuffer(byte_data))
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

func GetInventory() {
	fmt.Println("Please enter the product id")
	var product_id string
	_, err := fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err)
	}

	_, err = http.Get("http://localhost:8089/inventorydetail/" + fmt.Sprint(product_id))
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

func UpdateInventory() {
	fmt.Println("Please enter the product id")
	var product_id string
	_, err := fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err)
	}

	key := "quantity"

	fmt.Println("Please enter the quantity to be updated")
	var value int
	_, err = fmt.Scanln(&value)
	if err != nil {
		fmt.Println(err)
	}

	data := map[string]any{key: value}
	byte_data, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	request_body := bytes.NewBuffer(byte_data)
	req, err := http.NewRequest("PUT", fmt.Sprintf("http://localhost:8089/updateinventory"), request_body)
	if err != nil {
		fmt.Println(err)
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Update done succesfully")

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

func DeleteInventory() {
	fmt.Println("Please enter the product id")
	var product_id string
	_, err := fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8089/deleteinventory/%v", product_id), nil)
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
