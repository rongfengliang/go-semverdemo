package main

import (
	"log"

	"github.com/coreos/go-semver/semver"
)

func main() {
	first := semver.New("1.2.3")
	second := semver.New("1.2.4")
	log.Println(first.LessThan(*second))

}
