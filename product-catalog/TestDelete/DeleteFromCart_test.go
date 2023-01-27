package TestDelete

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestDeleteCartNotExists(t *testing.T) {
	data := []byte(`{"product_id":2, "ref":"axxyy"}`)

	req, err := http.NewRequest("DELETE", "http://localhost:8089/deletefromcart", bytes.NewBuffer(data))
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

	expected := "{\"message\":\"Product is not found in your cart\"}\n"

	bodyBytes, err := io.ReadAll(resp.Body)

	if string(bodyBytes) != expected {
		t.Errorf("unexpected: got %s, want %s", string(bodyBytes), expected)
	}

}

func TestDeleteCartExists(t *testing.T) {

	data := []byte(`{"product_id":7, "reference_id":"c6ac00da-f4fe-4941-8b20-c1a1f504612a"}`)

	req, err := http.NewRequest("DELETE", "http://localhost:8089/deletefromcart", bytes.NewBuffer(data))
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
	expected := "{\"message\":\"Product is not found in your cart\"}\n"

	bodyBytes, err := io.ReadAll(resp.Body)

	if string(bodyBytes) != expected {
		t.Errorf("unexpected: got %s, want %s", string(bodyBytes), expected)
	}

}
