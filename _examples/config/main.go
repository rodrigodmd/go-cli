package main

import (
	"fmt"
	"github.com/rodrigodmd/go-cli/cli"
)

func main() {
	cmd := cli.New("myapp", "1.0.1", )

	cmd.SetDefault(func(args []string) {
		fmt.Println("Default run!!")
	})

	cmd.AddRoute("get", "will get the thing inside the thing", func(args []string) {
		fmt.Println("Get Command!!")
		for _, arg := range args {
			fmt.Println(arg)
		}
	})

	cmd.Run()
}
