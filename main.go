package main

import (
	"fmt"

	"github.com/aqyuki/mcp/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("Error : %s\n", err)
	}
}
