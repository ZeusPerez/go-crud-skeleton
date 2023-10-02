package acceptance

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func makeRequest(t *testing.T, method, url string, body io.Reader) (int, string) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to make the request: %s", err)
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return resp.StatusCode, buf.String()
}
