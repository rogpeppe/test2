package main

import (
	"log"
	"runtime/debug"

	"github.com/kr/pretty"
)

func main() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		log.Println("no build info available")
		return
	}
	pretty.Println(info)
}
