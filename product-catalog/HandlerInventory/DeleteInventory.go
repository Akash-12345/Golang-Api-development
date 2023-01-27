package HandlerInventory

import (
	"net/http"
	"strconv"

	"github.com/akash-searce/product-catalog/DbConnect"
	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	"github.com/akash-searce/product-catalog/Response"
	response "github.com/akash-searce/product-catalog/Response"
	"github.com/gorilla/mux"
)

func DeleteInventory(w http.ResponseWriter, r *http.Request) {

	x := mux.Vars(r)["id"]
	db := DbConnect.ConnectToDB()
	a, _ := strconv.Atoi(x)
	if (a) <= 0 {
		Helpers.SendJResponse(Response.EnterValidInput, w)
		return
	}

	stmt, err := db.Prepare(queries.DeleteInventory)
	res, err := stmt.Exec(a)

	//	res, err := db.Exec(queries.DeleteInventory, x)

	if err == nil {

		count, err := res.RowsAffected()
		if err == nil {
			if count == 0 {
				Helpers.SendJResponse(response.ProductidDoesNotExist, w)
				return
			}
			Helpers.SendJResponse(response.InventoryDetailDeleted, w)
		}

	}

	return

}
