package edance

import (
	"fmt"
	"github.com/candbright/edance/module/utility"
	"os"
	"runtime"
	"strconv"
)

const (
	DbType = "mysql"
)

var (
	Port    = 10088
	DbIp    = ""
	DbPort  = ""
	LogFile = "/var/log/edance.log"
)

func InitGlobal() {
	initServicePort()
	if Port == 10088 {
		initLogFile("")
	} else {
		initLogFile("-" + strconv.Itoa(Port))
	}
	initDbHost()
}

func initServicePort() {
	if port := os.Getenv("EDANCE_PORT"); port != "" {
		portInt, err := strconv.Atoi(port)
		if err != nil || 0 < portInt || portInt > 65535 {
			return
		}
		Port = portInt
	}
}

func initLogFile(suffix string) {
	if runtime.GOOS == "linux" {
		LogFile = fmt.Sprintf("/var/log/edance%s.log", suffix)
	} else if runtime.GOOS == "windows" {
		dir, err := os.Getwd()
		if err != nil {
			LogFile = fmt.Sprintf("C:\\edance%s.log", suffix)
			return
		}
		LogFile = fmt.Sprintf("%s\\edance%s.log", dir, suffix)
	} else {
		LogFile = fmt.Sprintf("/var/log/edance%s.log", suffix)
	}
}

func initDbHost() {
	if ip := os.Getenv("EDANCE_DB_IP"); ip != "" {
		if utility.IsIpLegal(ip) {
			DbIp = ip
		}
	}
	if port := os.Getenv("EDANCE_DB_PORT"); port != "" {
		portInt, err := strconv.Atoi(port)
		if err != nil || 0 < portInt || portInt > 65535 {
			return
		}
		DbPort = port
	}
}
