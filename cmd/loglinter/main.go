package main

import (
	"log"

	"github.com/VeraVinogradova/loglinter/internal/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	a, err := analyzer.New()
	if err != nil {
		log.Fatal(err)
	}
	singlechecker.Main(a)
}
