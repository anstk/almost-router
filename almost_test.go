package almost

import (
	"io"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoute(t *testing.T) {

	index := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Index page")
	}

	user := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ok")
	}

	r := Router()
	r.Route("GET", "/", index)
	r.Route("GET", "/user", user)
	r.Route("POST", "/user", user)
	r.Route("MY", "/user", user)

	ts := httptest.NewServer(r)
	defer ts.Close()

	if _, body := testRequest(t, ts, "GET", "/", nil); body != "Index page" {
		t.Fatal(body)
	}

	if _, body := testRequest(t, ts, "GET", "/user", nil); body != "ok" {
		t.Fatal(body)
	}

	if _, body := testRequest(t, ts, "POST", "/user", nil); body != "ok" {
		t.Fatal(body)
	}

	if _, body := testRequest(t, ts, "MY", "/user", nil); body != "ok" {
		t.Fatal(body)
	}

}

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)

	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	return resp, string(respBody)
}
