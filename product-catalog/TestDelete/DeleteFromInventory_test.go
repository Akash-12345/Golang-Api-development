package TestDelete

import (
	"io"
	"net/http"
	"testing"
)

func TestDeleteInventoryNotExists(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:8089/deleteinventory/312", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	expected := "{\"message\":\"The value product_id does not exist, please enter a Valid ID\"}\n"

	bodyBytes, err := io.ReadAll(resp.Body)

	if string(bodyBytes) != expected {
		t.Errorf("unexpected: got %q, want %q", string(bodyBytes), expected)
	}

}

func TestDeleteInventoryExists(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:8089/deleteinventory/22", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
	}
	expected := "{\"message\":\"The inventory detail has been deleted successfully\"}\n"

	bodyBytes, err := io.ReadAll(resp.Body)

	if string(bodyBytes) != expected {
		t.Errorf("unexpected: got %q, want %q", string(bodyBytes), expected)
	}

}
