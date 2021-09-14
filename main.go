package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Invalid Argument")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == os.Args[1] {
			fmt.Printf("* %s\n", scanner.Text())
		} else {
			fmt.Printf("  %s\n", scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
