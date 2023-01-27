package TestInsert

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/akash-searce/product-catalog/typedefs"
)

func TestInsertCart(t *testing.T) {

	resp, err := http.Post("http://localhost:8089/addtocart?ref=c6ac00da-f4fe-4941-8b20-c1a1f504612a&product_id=2&quantity=1", "application/json", nil)
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	// Read the response from the API
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	response := typedefs.JResponse{}
	err = json.Unmarshal(body, &response)

	// Make assertions about the output of the function
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	if response.Message != "Respective products have been successfully added to the cart" {
		t.Errorf("Expected response body 'Respective products have been successfully added to the cart', got '%s'", response.Message)
	}

}

func TestInsertReferenceNotFound(t *testing.T) {
	resp, err := http.Post("http://localhost:8089/addtocart?ref=c6da-f4fe-4941-8b20-c1a1f504612a&product_id=7&quantity=8", "application/json", nil)
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	response := typedefs.JResponse{}
	err = json.Unmarshal(body, &response)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	if response.Message != "Invalid cart_reference" {
		t.Errorf("Expected response body 'Invalid cart_reference', got '%s'", response.Message)
	}
}

func TestInsertNotCart(t *testing.T) {

	resp, err := http.Post("http://localhost:8089/addtocart?ref=c6ac00da-f4fe-4941-8b20-c1a1f504612a&product_id=7&quantity=2000", "application/json", nil)
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	response := typedefs.JResponse{}
	err = json.Unmarshal(body, &response)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	if response.Message != "Inventory Quantity is less than the required quantity: 792" {
		t.Errorf("Expected 'Inventory Quantity is less than the required quantity: 792' got '%s'", response.Message)
	}

}
