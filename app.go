package main

import (
  "fmt"
)

func run(args []string) {
	for _, v := range args[1:] {
		fmt.Println(v)
	}
}
