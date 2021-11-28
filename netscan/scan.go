package netscan

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func scanPort(protocol string, hostname string, port int, results *[]scanResult, wg *sync.WaitGroup) {
	result := scanResult{port: strconv.Itoa(port), protocol: protocol}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 200*time.Millisecond)
	defer wg.Done()
	if err != nil {
		result.state = "Closed"
	} else {
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(conn)
		result.state = "Open"
		result.service = getService(strconv.Itoa(port), protocol)
		*results = append(*results, result)
	}
}

func getService(port string, protocol string) string {
	csvFile, err := os.Open("netscan/service-names-port-numbers.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			return "No service name found"
		} else if err != nil {
			log.Fatal(err)
		} else if line[1] == port && line[2] == protocol {
			return line[0]
		} else {
			portCsv, _ := strconv.Atoi(line[1])
			portArg, _ := strconv.Atoi(port)
			if portCsv > portArg {
				return "No service name found"
			}
		}
	}
}

func initialScan(arguments filterArg, results *[]scanResult) {
	srcIp, sPort := localIPPort(net.ParseIP(arguments.ip))
	if begin, err := strconv.Atoi(arguments.ports[0]); err == nil {
		if end, err := strconv.Atoi(arguments.ports[1]); err == nil {
			var wg sync.WaitGroup
			for i := begin; i <= end; i++ {
				if end-i <= arguments.speedup {
					wg.Add(end - i + 1)
					for j := i; j <= end; j++ {
						if arguments.scan == "syn" {
							go synScan(arguments.ip, j, results, &wg, srcIp, sPort)
						} else {
							go scanPort(arguments.scan, arguments.ip, j, results, &wg)
						}
					}
					i = end + 1
					wg.Wait()
				} else {
					wg.Add(arguments.speedup)
					for j := i; j < i+arguments.speedup; j++ {
						if arguments.scan == "syn" {
							go synScan(arguments.ip, j, results, &wg, srcIp, sPort)
						} else {
							go scanPort(arguments.scan, arguments.ip, j, results, &wg)
						}
					}
					i += arguments.speedup
					wg.Wait()
				}
			}
		}
		printResult(*results)
	}
}
