package TestInsert

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/akash-searce/product-catalog/typedefs"
)

func TestInsertCategory(t *testing.T) {
	data := []byte(`{"category_id":19, "category_name":"appliances"}`)

	resp, err := http.Post("http://localhost:8089/addcategory", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	response := typedefs.JResponse{}
	err = json.Unmarshal(body, &response)
	println(resp)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
	if response.Message != "Category detail has been added successfully" {
		t.Errorf("Expected Category detail has been added successfully got %s", response.Message)
	}

}
