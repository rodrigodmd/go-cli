package main

import (
	"fmt"
	"github.com/rodrigodmd/go-cli/cli"
)

func main() {
	cli.New("myapp", "1.0.0", func(args []string){
		fmt.Println("Hello World!!")
		for _, arg := range args {
			fmt.Println(arg)
		}
	}).Run()
}
