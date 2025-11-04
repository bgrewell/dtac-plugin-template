package main

import (
	"log"

	"github.com/bgrewell/dtac-agent/pkg/plugins"
	"{{MODULE_PATH}}/pkg/{{plugin}}"
)

func main() {
	p := {{plugin}}.New() // constructor
	h, err := plugins.NewPluginHost(p)
	if err != nil {
		log.Fatal(err)
	}
	if err := h.Serve(); err != nil {
		log.Fatal(err)
	}
}
