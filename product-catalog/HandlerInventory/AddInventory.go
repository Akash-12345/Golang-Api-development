package HandlerInventory

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/akash-searce/product-catalog/DbConnect"
	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	response "github.com/akash-searce/product-catalog/Response"
	"github.com/akash-searce/product-catalog/typedefs"
)

func AddIntoInventory(w http.ResponseWriter, r *http.Request) {
	var inventory = typedefs.Inventory{}
	reqBody, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(reqBody, &inventory)
	if err != nil {
		Helpers.SendJResponse(response.UnmarshalError, w)
	}
	if inventory.Product_Id < 0 || inventory.Quantity < 0 {
		Helpers.SendJResponse(response.EnterValidInput, w)
	}
	db := DbConnect.ConnectToDB()
	result, err := db.Exec("select * from inventory where product_id =$1;", inventory.Product_Id)
	if err != nil {
		Helpers.SendJResponse(response.ErrorInQuery, w)
		print(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		Helpers.SendJResponse(response.RowScanError, w)
		print(err)
	}
	if rows == 1 {
		Helpers.SendJResponse(response.InventoryProductAlreadyPresent, w)
		return
	}

	stmt, err := db.Prepare(queries.AddInventory)
	_, err = stmt.Exec(inventory.Product_Id, inventory.Quantity)
	if err != nil {
		Helpers.SendJResponse(response.ErrorInQuery, w)
	} else {
		Helpers.SendJResponse(response.InventoryDetailAdded, w)
	}

	/*

		var inventory typedefs.Inventory
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Error")
			Helpers.HandleError(err)
		}
		err = json.Unmarshal(reqBody, &inventory)
		if err != nil {
			Helpers.SendJResponse(Response.UnmarshalError, w)
			return
		}
		if inventory.Product_Id <= 0 || inventory.Quantity <= 0 {
			Helpers.SendJResponse(Response.EnterValidInput, w)
			return
		}
		db := DbConnect.ConnectToDB()
		stmt, err := db.Prepare(queries.AddInventory)
		_, err = stmt.Exec(inventory.Product_Id, inventory.Quantity)
		fmt.Println(inventory)

		if err != nil {
			Helpers.HandleError(err)
			fmt.Println(err) //check here

		} else {
			Helpers.SendJResponse(response.InventoryDetailAdded, w)
		}
	*/

}
