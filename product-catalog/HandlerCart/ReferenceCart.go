package HandlerCart

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	"github.com/akash-searce/product-catalog/Response"
	"github.com/google/uuid"
)

func CreateCart(w http.ResponseWriter, r *http.Request) {
	//create a cart and use specific cart reference id to access the cart.
	ref := uuid.New()

	_, err := Helpers.QueryRun(queries.InsertCartReference, ref, time.Now())
	if err != nil {
		Helpers.SendJResponse(Response.RunQueryError, w)
		fmt.Println(err)
		Helpers.HandleError(err)
	}
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		Helpers.HandleError(err)
	} else {
		json.NewEncoder(w).Encode((map[string]uuid.UUID{"ref": ref}))
	}
}
