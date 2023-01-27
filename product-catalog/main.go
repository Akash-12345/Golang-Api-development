package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akash-searce/product-catalog/HandlerCart"
	"github.com/akash-searce/product-catalog/HandlerCategory"
	"github.com/akash-searce/product-catalog/HandlerInventory"
	"github.com/akash-searce/product-catalog/HandlerProduct"
	"github.com/akash-searce/product-catalog/TConsoleInterface"

	//"github.com/akash-searce/product-catalog/t_console_interface"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	f := func() {
		fmt.Println("If you want to start console interface? (yes or no)")
		var reply string
		_, err := fmt.Scanln(&reply)
		if err != nil {
			fmt.Println("error in reading input!!")
		}
		if reply == "yes" {
			TConsoleInterface.Console()
		} else if reply == "no" {
			fmt.Println("Console Interface cancelled")
		}
	}

	go f()

	r := mux.NewRouter()
	// handlers for product_master table
	r.HandleFunc("/addproduct", HandlerProduct.AddProduct).Methods("POST")
	r.HandleFunc("/product/{id}", HandlerProduct.GetProduct).Methods("GET")
	r.HandleFunc("/getproducts", HandlerProduct.GetProducts).Methods("GET")
	r.HandleFunc("/updateproduct", HandlerProduct.UpdateProduct).Methods("PUT")
	r.HandleFunc("/deleteproduct/{id}", HandlerProduct.DeleteProduct).Methods("DELETE")
	// handlers for category_master table
	r.HandleFunc("/addcategory", HandlerCategory.AddCategory).Methods("POST")
	r.HandleFunc("/getcategory/{id}", HandlerCategory.GetCategory).Methods("GET")
	r.HandleFunc("/updatecategory", HandlerCategory.UpdateCategory).Methods("PUT")
	r.HandleFunc("/deletecategory/{id}", HandlerCategory.DeleteCategory).Methods("DELETE")
	//handlers for inventory table
	r.HandleFunc("/addinventory", HandlerInventory.AddIntoInventory).Methods("POST")
	r.HandleFunc("/inventorydetail/{id}", HandlerInventory.GetInventory).Methods("GET")
	r.HandleFunc("/updateinventory", HandlerInventory.UpdateInventory).Methods("PUT")
	r.HandleFunc("/deleteinventory/{id}", HandlerInventory.DeleteInventory).Methods("DELETE")
	// handlers for cart table
	r.HandleFunc("/cart/createreference", HandlerCart.CreateCart).Methods("POST")
	r.HandleFunc("/addtocart", HandlerCart.AddToCart).Methods("POST")
	r.HandleFunc("/cartitems/add", HandlerCart.AddItemsToCart).Methods("POST")
	r.HandleFunc("/cart/get", HandlerCart.GetCart1).Methods("GET")
	r.HandleFunc("/deletefromcart", HandlerCart.RemoveItemFromCart).Methods("DELETE")

	fmt.Printf("Starting server at port 8089\n")
	log.Fatal(http.ListenAndServe(":8089", r))
}
