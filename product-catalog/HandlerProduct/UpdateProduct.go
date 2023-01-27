package HandlerProduct

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/akash-searce/product-catalog/DbConnect"
	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	response "github.com/akash-searce/product-catalog/Response"
	"github.com/akash-searce/product-catalog/typedefs"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db := DbConnect.ConnectToDB()
	var newproduct typedefs.Product_master
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to insert")
		Helpers.HandleError(err)
	}

	err = json.Unmarshal(reqBody, &newproduct) //unmarshal the response and store into struct newproduct
	fmt.Println(newproduct)
	rows, err := db.Query("SELECT * from product_master WHERE product_id = $1", newproduct.Product_Id)
	if err != nil {
		fmt.Println("error while selecting product")
		Helpers.HandleError(err)
	}
	defer rows.Close()
	var existing_product typedefs.Product_master
	var rawContent []byte
	for rows.Next() {
		fmt.Println("working")
		err := rows.Scan(&existing_product.Product_Id, &existing_product.Name, &rawContent, &existing_product.SKU, &existing_product.Category_Id, &existing_product.Price)
		if err != nil {
			Helpers.HandleError(err)
			fmt.Println("error while scanning")
		}

		err = json.Unmarshal(rawContent, &existing_product.Specification)
		if err != nil {
			Helpers.HandleError(err)
			fmt.Println("error while unmarshalling")
		}
		if newproduct.Name == "" {
			newproduct.Name = existing_product.Name
		}
		if newproduct.Price == 0 {
			newproduct.Price = existing_product.Price
		}
		if newproduct.SKU == "" {
			newproduct.SKU = existing_product.SKU
		}
		if newproduct.Specification == nil {
			newproduct.Specification = existing_product.Specification
		}
		if newproduct.Category_Id == 0 {
			newproduct.Category_Id = existing_product.Category_Id
		}
		if newproduct.Category_Id != existing_product.Category_Id {
			result := fmt.Sprintf("Cannot alter the category id for product!!!\n please update in category table")
			json.NewEncoder(w).Encode(result)
			return
		}

		json_specification, err := json.Marshal(newproduct.Specification)

		fmt.Println(newproduct)
		db.Exec(queries.UpdateProduct, newproduct.Name, newproduct.SKU, newproduct.Price, json_specification, newproduct.Product_Id)
		if err != nil {
			Helpers.HandleError(err)
			fmt.Println("error", err)
		} else {
			Helpers.SendJResponse(response.ProductDetailUpdated, w)
		}
	}
}
