package Helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/akash-searce/product-catalog/typedefs"
)

func SendResponse(v any, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func SendJResponse(message string, w http.ResponseWriter) {
	SendResponse(typedefs.JResponse{Message: message}, w)
}

func HandleError(err error) {
	if err != nil {
		output := fmt.Sprint(err)
		//fmt.Println(output)

		file, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		log.SetOutput(file)
		log.Println(output)
	}
}
