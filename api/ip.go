package handler

import (
	"apis/lib"
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

	lib.WriteTextResponse(w, r, ip)
}
