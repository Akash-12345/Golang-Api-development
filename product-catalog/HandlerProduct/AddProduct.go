package HandlerProduct

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/akash-searce/product-catalog/DbConnect"
	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	"github.com/akash-searce/product-catalog/Response"
	response "github.com/akash-searce/product-catalog/Response"
	"github.com/akash-searce/product-catalog/typedefs"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var product typedefs.Product_master = typedefs.Product_master{}
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &product)
	if err != nil {
		Helpers.SendJResponse(Response.UnmarshalError, w)
		Helpers.HandleError(err)
		return
	}
	if product.Product_Id < 0 {
		Helpers.SendJResponse(Response.EnterValidInput, w)
		return
	}
	if product.Category_Id < 0 {
		Helpers.SendJResponse(Response.EnterValidInput, w)
		return
	}
	if product.Price < 0 {
		Helpers.SendJResponse(Response.EnterValidInput, w)
		return
	}
	fmt.Println(product.Product_Id)
	db := DbConnect.ConnectToDB()
	stmt, err := db.Prepare(queries.AddProduct)
	spec_json_byte, err := json.Marshal(product.Specification)
	_, err = stmt.Exec(product.Product_Id, product.Name, spec_json_byte, product.SKU, product.Category_Id, product.Price)

	if err != nil {
		Helpers.SendJResponse(Response.RunQueryError, w)
		fmt.Println(err) //check here
		Helpers.HandleError(err)
		return
	} else {
		Helpers.SendJResponse(response.ProductDetailAdded, w)
		return
	}
}
