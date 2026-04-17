package handler

import (
	"net"
	"net/http"
	"strings"
)

func getIP(r *http.Request) string {
	if ip := r.Header.Get("CF-Connecting-IP"); ip != "" {
		return ip
	}

	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		parts := strings.Split(ip, ",")
		return strings.TrimSpace(parts[0])
	}

	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	return host
}

func IPHandler(w http.ResponseWriter, r *http.Request) {
	ip := getIP(r)

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
	_, _ = w.Write([]byte(ip + "\n"))
}
