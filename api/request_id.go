package handler

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strings"
)

func generateRequestID() string {
	b := make([]byte, 12)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func FromRequest(r *http.Request) string {
	headers := []string{
		"X-Request-Id",
		"X-Correlation-Id",
	}

	for _, h := range headers {
		if v := strings.TrimSpace(r.Header.Get(h)); v != "" {
			return v
		}
	}

	return generateRequestID()
}
