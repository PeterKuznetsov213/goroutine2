package main

import (
	"fmt"
	"strings"
)

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all", "goodbye all", "hello again"} {
		downstream <- v
	}
	downstream <- ""
}

func filterGopher(upstream, downstream chan string) {
	itemCheck := ""
	for {
		item := <-upstream
		if item == "" {
			downstream <- ""
			return
		}
		if !strings.Contains(item, "bad") {
			if itemCheck != item {
				downstream <- item
			}
		}
		itemCheck = item
	}
}

func printGopher(upstream chan string) {
	for {
		v := <-upstream
		if v == "" {
			return
		}
		fmt.Println(v)
	}
}
