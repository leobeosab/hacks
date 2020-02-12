package helpers

import (
	"fmt"
	"net"
	"time"
)

func ReturnActiveWebPortURLS(host string) []string {
	portsToCheck := []int{80, 443, 8080, 8000, 3000}
	validHTTPServers := make([]string, 0)
	ssl := "https://"
	nonssl := "http://"

	// more bad design
	for _, port := range portsToCheck {
		if IsHTTPServerRunning(host, port) {
			var url string
			if port == 443 {
				url = ssl + host
			} else {
				url = fmt.Sprintf("%s%s:%d", nonssl, host, port)
			}

			validHTTPServers = append(validHTTPServers, url)
		}
	}

	return validHTTPServers
}

func IsHTTPServerRunning(host string, port int) bool {
	timeoutSeconds := 1
	host = fmt.Sprintf("%s:%d", host, port)
	_, err := net.DialTimeout("tcp", host, time.Duration(timeoutSeconds)*time.Second)

	return err == nil
}
