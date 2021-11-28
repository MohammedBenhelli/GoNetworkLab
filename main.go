package main

import (
	"github.com/MohammedBenhelli/GoNetworkLab/netscan"
	"github.com/akamensky/argparse"
	"os"
)

func main() {
	parser := argparse.NewParser("NetworkLab", "Some Go Network tools")
	tool := parser.String("t", "tool", &argparse.Options{Required: true, Help: "The tool name"})

	_ = parser.Parse(os.Args)
	if len(*tool) == 0 {
		panic(parser.Usage(""))
	}

	switch *tool {
	case "netscan":
		netscan.Main()
	default:
		netscan.Main()
	}
}
