package utility

import "net"

func GetLocalIp() string {
	return "127.0.0.1"
}

func IsIpLegal(ip string) bool {
	ipParse := net.ParseIP(ip)
	return ipParse != nil
}
