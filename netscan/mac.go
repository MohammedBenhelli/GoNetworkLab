package netscan

import (
	"fmt"
	"log"
	"net"
)

func getMacAddr() ([]macResult, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var results []macResult
	for _, ifa := range ifaces {
		address, _ := ifa.Addrs()
		for _, addr := range address {
			switch v := addr.(type) {
				case *net.IPNet:
					a := macResult{ifa.HardwareAddr.String(), v.IP, v.IP.DefaultMask().String()}
					if a._MAC != "" {
						results = append(results, a)
					}
				case *net.IPAddr:
					a := macResult{ifa.HardwareAddr.String(), v.IP, v.IP.DefaultMask().String()}
					if a._MAC != "" {
						results = append(results, a)
					}
			}
		}

	}
	return results, nil
}

func localIp()  {
	as, err := getMacAddr()
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range as {
		fmt.Println(a)
	}
}
