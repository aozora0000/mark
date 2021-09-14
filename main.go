package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var mark = flag.String("mark", "*", "current mark char")
var through = flag.Bool("through", true, "through empty line")

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("Invalid Argument")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if *through && scanner.Text() == "" {
			continue
		}
		if ArgMatch(flag.Args(), strings.TrimSpace(scanner.Text())) {
			fmt.Printf("%s %s\n", *mark, strings.TrimSpace(scanner.Text()))
		} else {
			fmt.Printf("%s %s\n", strings.Repeat(" ", len(*mark)), strings.TrimSpace(scanner.Text()))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
