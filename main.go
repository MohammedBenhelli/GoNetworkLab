package main

import (
	"fmt"
	"github.com/MohammedBenhelli/GoNetworkLab/arg"
	"github.com/MohammedBenhelli/GoNetworkLab/scanner"
)

var (
	results []scanner.ScanResult
)

func main() {
	arguments := arg.Arg()
	fmt.Printf("%+v\n", arguments)
	if !arguments.File() {
		scanner.InitialScan(arguments, &results)
	} else {
		scanner.FileScan(arguments, &results)
	}
	//scanner.LocalIp()
	//ifs, _ := pcap.FindAllDevs()
	//for i := range ifs {
	//	fmt.Println(ifs[i])
	//}
}
