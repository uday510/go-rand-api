package handler

import (
	"net"
	"net/http"
	"strings"
)

func getIP(r *http.Request) string {
	ip := r.Header.Get("CF-Connecting-IP")
	if ip != "" {
		return ip
	}

	ip = r.Header.Get("X-Forwarded-For")
	if ip != "" {
		parts := strings.Split(ip, ",")
		return strings.TrimSpace(parts[0])
	}

	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
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

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(ip + "\n"))
}
