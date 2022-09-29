package main

import (
	"fmt"
	"os"

	"github.com/casbin/casbin/v2"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <user> <resource> <action>\n", os.Args[0])
		os.Exit(1)
	}

	e, _ := casbin.NewEnforcer("model.conf", "policy.csv")
	res, _ := e.Enforce(os.Args[1], os.Args[2], os.Args[3])
	if res {
		fmt.Println("Access allowed")
	} else {
		fmt.Println("Access denied")
	}
}
