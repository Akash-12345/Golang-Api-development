package HandlerProduct

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

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]
	db := DbConnect.ConnectToDB()
	a, _ := strconv.Atoi(x)
	if (a) < 0 {
		Helpers.SendJResponse(Response.EnterValidInput, w)
		return
	}

	res, err := db.Exec(queries.DeleteProduct, x)

	if err == nil {

		count, err := res.RowsAffected() // by default rows affected would return count and error.
		if err == nil {
			if count == 0 {
				Helpers.SendJResponse(response.ProductidDoesNotExist, w)
				return // if code hits return then it wont run the code written after it
			}
			Helpers.SendJResponse(response.ProductDetailDeleted, w)
		}

	}

	return

}

/*Params := mux.Vars(r)
	db := dbconnect.ConnectToDB()
	ID := Params["id"]
	stmt, err := db.Prepare("DELETE FROM product_master where product_id=$1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stmt)
	if err != nil {
		fmt.Fprintf(w, "Error")
	} else {
		fmt.Println("successfully deleted ")
	}
	_, err = stmt.Query(ID)
	defer stmt.Close()

	if err != nil {
		fmt.Fprintf(w, "Error")
	}

}
*/
