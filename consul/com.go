package consul

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
)

var count int64

func LocalIP() string {
	adders, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range adders {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func ParseIp(ip string) string {
	address := net.ParseIP(ip)
	if address == nil {
		return ""
	}
	return address.String()
}

// Check consul 服务端会自己发送请求，来进行健康检查
func Check(w http.ResponseWriter, r *http.Request) {
	s := "Check" + fmt.Sprint(count) + "remote:" + r.RemoteAddr + " " + r.URL.String()
	fmt.Println(s)
	fmt.Fprintln(w, s)
	count++
}

func CheckIPAddr(ip string) error {
	if ip == "" {
		return errors.New(fmt.Sprintf("Address error : %v", "注册的服务IP不能为空"))
	}
	if len(strings.Split(ip, ":")) < 1 {
		return errors.New(fmt.Sprintf("Address error : %v", "注册的服务IP格式必须是 ip:port "))
	}
	if ParseIp(strings.Split(ip, ":")[0]) == "" && strings.ToLower(strings.Split(ip, ":")[0]) != "localhost" {
		return errors.New(fmt.Sprintf("Address error : %v%v", ip, ":注册的服务IP格式不正确"))
	}
	return nil
}
