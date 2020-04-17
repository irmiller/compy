package transcoder

import (
	"net/http"
	"strings"
	"log"
)

func SupportsWebP(headers http.Header) bool {
	log.Printf(logtest)
	for _, v := range strings.Split(headers.Get("Accept"), ",") {
		if strings.SplitN(v, ";", 2)[0] == "image/webp" {
			return true
		}
	}
	return false
}
