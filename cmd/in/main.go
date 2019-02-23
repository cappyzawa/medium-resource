package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cappyzawa/medium-resource/in"
)

func main() {
	var req in.Request
	if err := json.NewDecoder(os.Stdin).Decode(&req); err != nil {
		fmt.Fprintf(os.Stderr, "invalid payload: %s\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "{\"version\":{\"id\": \"%s\"}}", req.Version.ID)
}
