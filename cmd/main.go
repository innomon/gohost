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
	builtins(vm.Globals)

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

func builtins(dict starlark.StringDict) {
	dict["http"] = nethttp.NewModule()
	// TODO add the modules as needed from  https://github.com/qri-io/starlib
	// TODO add Args[] as builtin
	// TODO add the modules as needed from  https://github.com/1set/starlet/blob/master/config.go#L65
	// TODO print builtins
}
