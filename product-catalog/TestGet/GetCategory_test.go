package TestGet

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/akash-searce/product-catalog/typedefs"
)

func TestGetCategory(t *testing.T) {

	resp, err := http.Get("http://localhost:8089/getcategory/3")
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

}

func TestGetCategoryNotExist(t *testing.T) {
	// Make a request to the API endpoint that triggers the insert function
	resp, err := http.Get("http://localhost:8089/getcategory/222")
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	// Read the response from the API
	body, err := ioutil.ReadAll(resp.Body)
	response := typedefs.JResponse{}

	err = json.Unmarshal(body, &response)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	// Make assertions about the output of the function
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	if response.Message != "Category id not present" {
		t.Errorf("Expected Category id not present, got %s", response.Message)
	}

}
