package main

import (
	"fmt"

	"github.com/magiconair/properties"
)

func main() {
	fmt.Println("Hello, World")
	p := properties.MustLoadFile("sim.properties",
		properties.UTF8)
	if port, ok := p.Get("port"); ok {
		fmt.Println(port)
	} else {
		fmt.Println("Port not found")
	}
}
