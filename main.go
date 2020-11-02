package main

import (
	"fmt"
	"net/http"
	"os"
)

func must(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to %s: %v\n", msg, err)
		panic(err)
	}
}

func main() {
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	must("serve webpages", err)
}
