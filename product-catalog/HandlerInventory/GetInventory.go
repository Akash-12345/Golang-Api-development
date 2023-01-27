package HandlerInventory

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/akash-searce/product-catalog/DbConnect"
	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	"github.com/akash-searce/product-catalog/Response"
	response "github.com/akash-searce/product-catalog/Response"
	"github.com/akash-searce/product-catalog/typedefs"
	"github.com/gorilla/mux"
)

func GetInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	inventory := typedefs.Inventory{}
	db := DbConnect.ConnectToDB()
	a, _ := strconv.Atoi(ID)
	if a <= 0 {
		Helpers.SendJResponse(Response.EnterValidInput, w)
		return
	}
	stmt, err := db.Prepare(queries.GetInventory)
	if err != nil {
		Helpers.SendJResponse(Response.RunQueryError, w)
		Helpers.HandleError(err)
		fmt.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(ID)
	if err != nil {
		Helpers.SendJResponse(Response.RunQueryError, w)
		Helpers.HandleError(err)
		fmt.Println(err)
	}
	if rows.Next() {
		rows.Scan(&inventory.Product_Id, &inventory.Quantity)
		json.NewEncoder(w).Encode(inventory)
		fmt.Println(inventory)
	} else {
		Helpers.SendJResponse(response.ProductNotFound, w)
	}

	/*
		Params := mux.Vars(r)
		inventory := typedefs.Inventory{}
		db := dbconnect.ConnectToDB()
		ID := Params["id"]

		stmt, err := db.Prepare("SELECT * from inventory where product_id=$1")
		if err != nil {
			fmt.Println(err)
		}
		defer stmt.Close()
		rows, err := stmt.Query(ID)
		if err != nil {
			panic(err)
		}

		if rows.Next() {
			rows.Scan(&inventory.Product_Id, &inventory.Quantity)
			json.NewEncoder(w).Encode(inventory)
		} else {
			json.NewEncoder(w).Encode(map[string]string{"message": "Product does not exist"})
		}
	*/
}
