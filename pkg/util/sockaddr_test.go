package util_test

import (
	"net"
	"testing"

	"github.com/nielsAD/noot/pkg/util"
)

func TestSockAddrConv(t *testing.T) {
	ipA, e := net.ResolveIPAddr("ip", "127.0.0.1")
	if e != nil {
		t.Fatal(e)
	}
	ipB := util.Addr(ipA)
	if ipA.String() != ipB.IPAddr().String() {
		t.Fatal("ResolveIPAddr != IpAddr")
	}

	udpA, e := net.ResolveUDPAddr("udp", "127.0.0.1:6112")
	if e != nil {
		t.Fatal(e)
	}
	udpB := util.Addr(udpA)
	if udpA.String() != udpB.UDPAddr().String() {
		t.Fatal("ResolveUDPAddr != UDPAddr")
	}

	tcpA, e := net.ResolveTCPAddr("tcp", "127.0.0.1:6112")
	if e != nil {
		t.Fatal(e)
	}
	tcpB := util.Addr(tcpA)
	if tcpA.String() != tcpB.TCPAddr().String() {
		t.Fatal("ResolveTCPAddr != TCPAddr")
	}
}
