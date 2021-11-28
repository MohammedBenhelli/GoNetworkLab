package netscan

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
	"strconv"
	"strings"
)

func filterPorts(str string) []string {
	num := 0
	if strings.Contains(str, "-") {
		ret := strings.Split(str, "-")
		for i := 0; i < len(ret); i += 2 {
			if i+2 < len(ret) || i == 0 {
				sum1, _ := strconv.Atoi(ret[i+1])
				sum2, _ := strconv.Atoi(ret[i])
				num += sum1 - sum2
			} else {
				num++
			}
		}
		fmt.Printf("\n\nNo of Ports to scan : %s\n\n", strconv.Itoa(num))
		return ret
	} else {
		fmt.Printf("\n\nNo of Ports to scan : 1\n\n")
		return []string{str}
	}
}

func Main() {
	// -i 127.0.0.1 -p 80-8000 -u 1000 -s tcp
	parser := argparse.NewParser("NetScan", "NetScan tool")
	_ = parser.String("t", "type", &argparse.Options{Required: true, Help: "The type of tool"})
	ports := parser.String("p", "ports", &argparse.Options{Required: true, Help: "The ports to scan"})
	ip := parser.String("i", "ip", &argparse.Options{Required: false, Help: "The target ip to scan"})
	file := parser.String("f", "file", &argparse.Options{Required: false, Help: "The file containing the targets ips"})
	speedup := parser.String("u", "speedup", &argparse.Options{Required: true, Help: "The numbers of threads"})
	scan := parser.String("s", "scan", &argparse.Options{Required: true, Help: "The scan type"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(0)
	}

	speed, err := strconv.Atoi(*speedup)
	if err != nil {
		panic(err)
	}

	arguments := filterArg{
		filterPorts(*ports),
		*ip,
		speed,
		*scan,
		*file,
	}

	results := make([]scanResult, 0)
	if len(arguments.file) == 0 {
		initialScan(arguments, &results)
	} else {
		fileScan(arguments, &results)
	}
	//localIp()
	//ifs, _ := pcap.FindAllDevs()
	//for i := range ifs {
	//	fmt.Println(ifs[i])
	//}
}
