package HandlerCart

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
)

func RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()

	ref := urlQuery.Get("ref")
	product_id := urlQuery.Get("product_id")
	p_value, _ := strconv.Atoi(product_id)

	if p_value <= 0 {
		Helpers.SendJResponse(Response.EnterValidInput, w)
		return
	}
	if ref == "" || product_id == "" {
		Helpers.SendJResponse(response.ReferenceOrProductidMissing, w)
		return
	}

	rows, err := Helpers.QueryRun(queries.GetFromCartReference, ref)
	if err != nil {
		Helpers.SendJResponse(Response.RunQueryError, w)
		fmt.Println(err)
		Helpers.HandleError(err)
	}

	if !rows.Next() {
		Helpers.SendJResponse(response.InvalidCartRef, w)
		return
	}

	result, err := DbConnect.ConnectToDB().Exec(queries.DeleteFromCart, ref, product_id)
	if err != nil {
		Helpers.SendJResponse(Response.RunQueryError, w)
		fmt.Println(err)
		Helpers.HandleError(err)
	}

	rows_affected, err := result.RowsAffected()
	if err != nil {
		Helpers.SendJResponse(Response.RowsAffectedError, w)
		fmt.Println(err)
		Helpers.HandleError(err)
	}

	if rows_affected != 0 {
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
			Helpers.HandleError(err)
		} else {
			Helpers.SendJResponse(response.CartItemDeleted, w)
		}
	} else {
		Helpers.SendJResponse(response.ProductNotInCart, w)
	}
}
