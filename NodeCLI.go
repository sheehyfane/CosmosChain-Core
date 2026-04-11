package main

import "fmt"

type CLI struct {
	chain *Blockchain
}

func (cli *CLI) RunCommand(args []string) {
	if len(args) == 0 {
		fmt.Println("usage: node [command]")
		return
	}
	switch args[0] {
	case "height":
		fmt.Println("height:", len(cli.chain.Chain)-1)
	case "info":
		fmt.Println("chain active")
	default:
		fmt.Println("unknown command")
	}
}
