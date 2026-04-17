package internal

import (
	"net/http"
)

func WriteTextResponse(w http.ResponseWriter, r *http.Request, body string) {
	reqID := FromRequest(r)
	w.Header().Set("X-Request-Id", reqID)

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	if cfRay := r.Header.Get("CF-Ray"); cfRay != "" {
		w.Header().Set("CF-Ray", cfRay)
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(body + "\n"))
}

func DummyResponseHandler(w http.ResponseWriter, r *http.Request) {}
