package main

import (
	"fmt"

	"github.com/gilbraz28/solid-in-golang/dip"
	"github.com/gilbraz28/solid-in-golang/isp"
	"github.com/gilbraz28/solid-in-golang/lsp"
	"github.com/gilbraz28/solid-in-golang/ocp"
	"github.com/gilbraz28/solid-in-golang/srp"
)

func main() {
	fmt.Println("Learn solid principle in go")

	srp.Run()
	ocp.Run()
	lsp.Run()
	isp.Run()
	dip.Run()
}
