package main

import (
	"log"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

func main() {
	var config Config
	ctx := NewEvalContext()
	err := hclsimple.DecodeFile("testdata.hcl", ctx, &config)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	log.Printf("Configuration is %#v", config)
}
