package main

import "github.com/polarislabs/chamber/v2/cmd"

var (
	// This is updated by linker flags during build
	Version           = "dev"
	AnalyticsWriteKey = ""
)

func main() {
	cmd.Execute(Version, AnalyticsWriteKey)
}
