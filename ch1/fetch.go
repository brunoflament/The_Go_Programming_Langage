package main

import (
	"fmt"
	"os"
	"net/http"
)

func main() {
	for _, arg := range os.Args[1:] {
		resp, err := http.Get(arg)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}
		fmt.Print(*resp)
	}
}

