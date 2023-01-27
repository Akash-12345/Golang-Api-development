package TConsoleInterface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Category_master() {
	fmt.Println("Hi Welcome to our categories section please be free to perform CRUD operations on 'Category_master' table")
	fmt.Printf("1.Add\n2.Get\n3.Update\n4.Delete\n")
	fmt.Println("Please enter your choice")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(err)
	}
	if choice == 1 {
		AddCategory()
	} else if choice == 2 {
		GetCategory()
	} else if choice == 3 {
		UpdateCategory()
	} else if choice == 4 {
		DeleteCategory()
	}
}

func AddCategory() {
	fmt.Println("Please enter the valid category id")
	var category_id int
	_, err := fmt.Scanf("%d", &category_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the name for the category id")
	var category_name string
	_, err = fmt.Scanln(&category_name)
	if err != nil {
		fmt.Println(err)
	}

	own_data := fmt.Sprintf("{\"category_id\":%v,\"category_name\":\"%v\"}", category_id, category_name)

	byte_data := []byte(own_data)

	_, err = http.Post("http://localhost:8089/addcategory", "application/json", bytes.NewBuffer(byte_data))
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

func GetCategory() {
	fmt.Println("Please enter the category id")
	var category_id int
	_, err := fmt.Scanln(&category_id)
	if err != nil {
		fmt.Println(err)
	}

	_, err = http.Get("http://localhost:8089/getcategory/" + fmt.Sprint(category_id))
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

func UpdateCategory() {
	fmt.Println("Please enter the category id")
	var category_id int
	_, err := fmt.Scanln(&category_id)
	if err != nil {
		fmt.Println(err)
	}

	key := "category_name"

	fmt.Println("Please enter the name to be updated")
	var value string
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
	req, err := http.NewRequest("PUT", fmt.Sprintf("http://localhost:8089/updatecategory%v", category_id), request_body)
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

func DeleteCategory() {
	fmt.Println("Please enter the category id")
	var category_id int
	_, err := fmt.Scanln(&category_id)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8089/deletecategory/%v", category_id), nil)
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
