package HandlerCategory

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

func AddCategory(w http.ResponseWriter, r *http.Request) {
	println("hi")
	var category typedefs.Category_master = typedefs.Category_master{}
	reqBody, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(reqBody, &category)
	if err != nil {
		Helpers.SendJResponse(Response.UnmarshalError, w)
		Helpers.HandleError(err)
		return
	}
	fmt.Println(category.Category_Id, category.Category_Name)
	//unmarshal the json values from postman to put into database
	if category.Category_Id <= 0 {
		Helpers.SendJResponse(Response.EnterValidInput, w)
		return
	}
	db := DbConnect.ConnectToDB()
	stmt, err := db.Prepare(queries.AddCategory)
	if err != nil {
		Helpers.SendJResponse(Response.RunQueryError, w)
		Helpers.HandleError(err)
		return
	}
	_, err = stmt.Exec(category.Category_Id, category.Category_Name)

	if err != nil {
		Helpers.SendJResponse(Response.RunQueryError, w)
		fmt.Println(err)
		Helpers.HandleError(err)
		//check here
		return
	} else {
		Helpers.SendJResponse(response.CategoryDetailAdded, w)
		return
	}

}
