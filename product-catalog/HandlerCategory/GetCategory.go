package HandlerCategory

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	"github.com/akash-searce/product-catalog/Response"
	response "github.com/akash-searce/product-catalog/Response"
	"github.com/akash-searce/product-catalog/typedefs"
	"github.com/gorilla/mux"
)

func GetCategory(w http.ResponseWriter, r *http.Request) {
	/*Params := mux.Vars(r)
	db := dbconnect.ConnectToDB()
	ID := Params["id"]
	*/

	category := typedefs.Category_master{}
	getcategory_id := mux.Vars(r)["id"]
	x, _ := strconv.Atoi(getcategory_id)
	if x <= 0 {
		Helpers.SendJResponse(Response.EnterValidInput, w)
		return
	}

	rows, err := Helpers.QueryRun(queries.GetCategory, getcategory_id)
	if err != nil {
		Helpers.SendJResponse(Response.RunQueryError, w)
		Helpers.HandleError(err)
		fmt.Println(err)
	}
	if rows.Next() {
		err := rows.Scan(&category.Category_Id, &category.Category_Name)
		json.NewEncoder(w).Encode(category)
		fmt.Println(category)
		if err != nil {
			Helpers.SendJResponse(Response.RowScanError, w)
			Helpers.HandleError(err)
			fmt.Println(err)
		}
	} else {
		Helpers.SendJResponse(response.CategoryidNotPresent, w)
	}

}
