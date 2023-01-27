package HandlerCategory

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/akash-searce/product-catalog/DbConnect"
	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	"github.com/akash-searce/product-catalog/Response"
	response "github.com/akash-searce/product-catalog/Response"
	"github.com/gorilla/mux"
)

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]
	db := DbConnect.ConnectToDB()

	res, err := db.Exec(queries.DeleteCategory, x)

	a, _ := strconv.Atoi(x)
	if a <= 0 {
		Helpers.SendJResponse(Response.EnterValidInput, w)
		return
	}

	if err == nil {

		count, err := res.RowsAffected()
		if err == nil {
			if count == 0 {
				Helpers.SendJResponse(response.CategoryIdDoesNotExist, w)
				//result := fmt.Sprint("The value category_id does not exist,enter a Valid ID")
				//json.NewEncoder(w).Encode(result)
				return
			}
			json.NewEncoder(w).Encode(response.CategoryDetailDeleted)
		}

	}

	return

}
