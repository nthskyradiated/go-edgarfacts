package main

import (
	"fmt"
	"github.com/nthskyradiated/go-edgarfacts/internal/facts"
)

func main() {
cik := ""
name := ""
org := ""
email := ""

facts, err := facts.LoadFacts(cik, name, org, email)
if err != nil {
	panic(err)
}
fmt.Println(string(facts))
}