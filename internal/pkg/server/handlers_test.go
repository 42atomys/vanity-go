package server

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProxyHandler(t *testing.T) {
	LoadConfig() // nolint
	cachedTemplate = template.Must(template.ParseFiles("templates/index.html"))

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Host = "atomys.lab"

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(proxyHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	tests := []struct {
		name     string
		contains string
	}{
		{
			"vanity-go",
			"<meta name=\"go-import\" content=\"vanity-go git https://github.com/42Atomys/vanity-go.git\">",
		},
		{
			"dns-updater",
			"<meta name=\"go-import\" content=\"dns-updater git https://github.com/42Atomys/dns-updater.git\">",
		},
		{
			"subpath/gw2api-go",
			"<meta name=\"go-import\" content=\"subpath/gw2api-go git https://gitlab.com/Atomys/gw2api-go.git\">",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !strings.Contains(rr.Body.String(), tt.contains) {
				t.Errorf("index body can't contains valid meta body = %v, want %v", rr.Body.String(), tt.contains)
			}
		})
	}
}
