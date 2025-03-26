package fileops

import (
	"os"

	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

var Module = &starlarkstruct.Module{
	Name: "file",
	Members: starlark.StringDict{
		"load": starlark.NewBuiltin("load", load),
	},
}

func load(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var flname starlark.Value

	if err := starlark.UnpackPositionalArgs("load", args, kwargs, 1, &flname); err != nil {
		return nil, err
	}
	str, err := os.ReadFile(flname.String())
	if err != nil {
		return nil, err
	}
	return starlark.String(string(str)), nil
}
