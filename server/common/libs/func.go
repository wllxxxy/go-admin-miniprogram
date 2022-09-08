package libs

import (
	"net"
	"net/http"
)

func RemoteIp(r *http.Request) string {
	remoteIp := r.RemoteAddr

	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		remoteIp = ip
	} else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
		remoteIp = ip
	} else {
		remoteIp, _, _ = net.SplitHostPort(remoteIp)
	}

	//本地ip
	if remoteIp == "::1" {
		remoteIp = "127.0.0.1"
	}

	return remoteIp

}

//func ClientIP(r *http.Request) string {
//	xForwardedFor := r.Header.Get("X-Forwarded-For")
//	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
//	if ip != "" {
//		return ip
//	}
//	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
//	if ip != "" {
//		return ip
//	}
//	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
//		return ip
//	}
//	return ""
//}
