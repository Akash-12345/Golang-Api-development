package HandlerCart

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akash-searce/product-catalog/DbConnect"
	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	"github.com/akash-searce/product-catalog/Response"
	response "github.com/akash-searce/product-catalog/Response"
	"github.com/akash-searce/product-catalog/typedefs"
)

func GetCart1(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()
	reference := urlQuery.Get("ref")
	db := DbConnect.ConnectToDB()

	rows, err := db.Query(queries.JoinProductMasterCartItem, reference)
	if err != nil {
		Helpers.SendJResponse(Response.ErrorInQuery, w)
		fmt.Println(err)
		Helpers.HandleError(err)
	}
	if rows.Next() == false {
		w.Header().Add("Content-Type", "application/json")
		Helpers.SendJResponse(response.ReferenceIdNotExist, w)
	} else {
		list_of_cart := []typedefs.Newcart{}

		if err != nil {
			Helpers.SendJResponse(Response.ErrorInCategory, w)
			fmt.Println(err)
			Helpers.HandleError(err)
		}
		var total float32

		for rows.Next() {
			new_cart := typedefs.Newcart{}
			err := rows.Scan(&new_cart.Price, &new_cart.Product_name, &new_cart.Quantity)
			if err != nil {
				Helpers.SendJResponse(Response.ErrorInRowsNext, w)
				fmt.Println(err)
			}
			total += (float32(new_cart.Price) * float32(new_cart.Quantity))
			list_of_cart = append(list_of_cart, new_cart)
		}

		if len(list_of_cart) == 0 {
			Helpers.SendJResponse(response.DataNotFound, w)
			return
		}
		Cart := map[string]interface{}{
			"Total Price": total,
			"data":        list_of_cart,
		}
		json.NewEncoder(w).Encode(Cart)

	}
}
