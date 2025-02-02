package main

import (
	"context"
	_ "embed"
	"fmt"
	"gohost"
	"os"

	nethttp "github.com/pcj/starlark-go-nethttp"
	"go.starlark.net/starlark"
)

//go:embed default.star
var defaultConfig []byte
var err error

func main() {
	vm := &gohost.VmStarlark{Globals: make(starlark.StringDict)}
	vm.Globals["http"] = nethttp.NewModule()

	if len(os.Args) == 1 {
		err = vm.Exec(context.Background(), "default.star", defaultConfig)
	} else {
		err = vm.Exec(context.Background(), os.Args[1], nil)
	}
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print("OK")
	}

}
