package samples

import (
	"net"
	"sync"
	"testing"
)

var (
	service   map[string]net.Addr
	serviceMu sync.Mutex
)

func Test_Unprotected(t *testing.T) {
	WrongRegisterService("Lines", &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 0})
	WrongLockupService("Lines")

	RightRegisterService("Lines", &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 0})
	RightLockupService("Lines")
}

func WrongRegisterService(name string, addr net.Addr) {
	service[name] = addr
}

func WrongLockupService(name string) net.Addr {
	return service[name]
}

func RightRegisterService(name string, addr net.Addr) {
	serviceMu.Lock()
	defer serviceMu.Unlock()
	service[name] = addr
}

func RightLockupService(name string) net.Addr {
	serviceMu.Lock()
	defer serviceMu.Unlock()
	return service[name]
}
