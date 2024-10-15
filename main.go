package main

import (
	"os"

	"github.com/storskegg/utf-reporter/application/utfreporter"
)

func main() {
	if err := utfreporter.Execute(); err != nil {
		os.Exit(1)
	}
}
