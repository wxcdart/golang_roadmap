package httptestexamples

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Unit-style test: test handler directly using ResponseRecorder
func TestHelloHandler_Direct(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=Alice", nil)
	rr := httptest.NewRecorder()

	handler := SetupRouter()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if got := rr.Body.String(); got != "Hello, Alice" {
		t.Fatalf("unexpected body: %q", got)
	}
}

// Integration-style test: start httptest.Server and call via HTTP client
func TestEchoHandler_Server(t *testing.T) {
	ts := httptest.NewServer(SetupRouter())
	defer ts.Close()

	payload := []byte("payload-data")
	resp, err := http.Post(ts.URL+"/echo", "application/octet-stream", bytes.NewReader(payload))
	if err != nil {
		t.Fatalf("post error: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	if !bytes.Equal(body, payload) {
		t.Fatalf("unexpected echo body: %q", string(body))
	}
}
