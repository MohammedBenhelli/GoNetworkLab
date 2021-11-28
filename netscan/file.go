package netscan

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
)

func fileScan(arguments filterArg, results *[]scanResult) {
	ipList := fileToTab(arguments.file)
	for _, ip := range ipList {
		arguments.ip = ip
		color.Green("Ip tested %s\n", ip)
		initialScan(arguments, results)
		printResult(*results)
		*results = []scanResult{}
	}
}

func fileToTab(path string) []string {
	var result []string
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return []string{}
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			result = append(result, scanner.Text())
		}
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	return result
}

func printResult(results []scanResult) {
	for i := 0; i < len(results); i++ {
		fmt.Printf("%s/%s\t%s\t%s\n", results[i].port, results[i].protocol, results[i].state, results[i].service)
	}
}
