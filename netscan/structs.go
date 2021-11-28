package netscan

import "net"

type filterArg struct {
	ports   []string
	ip      string
	speedup int
	scan    string
	file    string
}

type scanResult struct {
	port     string
	state    string
	service  string
	protocol string
}

type macResult struct {
	_MAC  string
	_IP   net.IP
	_Mask string
}
