package main

import (
	"fmt"
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	ip4 := strings.Split(IP, ".")
	if len(ip4) == 4 && checkIPV4(ip4) {
		return "IPv4"
	}
	ip6 := strings.Split(IP, ":")
	if len(ip6) == 8 && checkIPV6(ip6) {
		return "IPv6"
	}
	return "Neither"
}

func checkIPV4(ips []string) bool {
	for _, v := range ips {
		if len(v) == 0 {
			return false
		} else if len(v) != 1 && v[0] == '0' {
			return false
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		if i > 255 || i < 0 {
			return false
		}
	}
	return true
}

func checkIPV6(ips []string) bool {
	for _, v := range ips {
		length := len(v)
		if length == 0 || length > 4 {
			return false
		}

		str := strings.ToLower(v)
		for i := 0; i < len(str); i++ {
			value := str[i]
			if (value >= '0' && value <= '9') || (value >= 'a' && value <= 'f') {
				continue
			}
			return false
		}
	}
	return true
}

func main() {
	addr := "172.16.254.1"
	fmt.Println(validIPAddress(addr))
	return
}
