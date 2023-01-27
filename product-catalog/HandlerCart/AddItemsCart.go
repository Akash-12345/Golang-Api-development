package HandlerCart

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akash-searce/product-catalog/Helpers"
	"github.com/akash-searce/product-catalog/Response"
)

func AddItemsToCart(w http.ResponseWriter, r *http.Request) {
	response := []map[string]any{}
	request_body := []map[string]int{}

	ref := r.URL.Query().Get("ref")

	err := json.NewDecoder(r.Body).Decode(&request_body)
	if err != nil {
		Helpers.SendJResponse(Response.ErrorInDecoding, w)
		Helpers.HandleError(err)
	}

	for _, v := range request_body {
		new_response_item := map[string]any{}
		product_id := v["product_id"]
		quantity := v["quantity"]

		url := "http://localhost:8079/addtocart?ref=" + ref + "&product_id=" + fmt.Sprint(product_id) + "&quantity=" + fmt.Sprint(quantity)
		//fmt.Println(url)
		_, err = http.Post(url, "application/json", nil)
		if err != nil {
			fmt.Println("ERROR IN DECODING", err)

		}

		new_response_item["product_id"] = product_id
		new_response_item["quantity"] = quantity
		new_response_item["response"] = "NEW PRODUCTS HAVE BEEN ADDED SUCCESFULLY"

		response = append(response, new_response_item)
	}

	json.NewEncoder(w).Encode(response)
}
