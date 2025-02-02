package gohost

import (
	"context"

	json_s "go.starlark.net/lib/json"
	math_s "go.starlark.net/lib/math"
	proto_s "go.starlark.net/lib/proto"
	time_s "go.starlark.net/lib/time"
	"go.starlark.net/repl"
	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
)

var Options = syntax.FileOptions{
	Set:               true,
	While:             true,
	TopLevelControl:   true,
	GlobalReassign:    true,
	LoadBindsGlobally: true,
	Recursion:         true,
}

func init() {
	// /starlark/library.go  : builtin modules
	starlark.Universe["json"] = json_s.Module
	starlark.Universe["math"] = math_s.Module
	starlark.Universe["proto"] = proto_s.Module
	starlark.Universe["time"] = time_s.Module

}

type VmStarlark struct {
	thread  *starlark.Thread
	Globals starlark.StringDict
}

func (vs *VmStarlark) Exec(ctx context.Context, file string, src interface{}) error {
	vs.thread = &starlark.Thread{Load: repl.MakeLoadOptions(&Options)}
	if vs.Globals == nil {
		vs.Globals = make(starlark.StringDict)
	}
	// set context
	vs.thread.SetLocal("context", ctx)
	_, err := starlark.ExecFileOptions(&Options, vs.thread, file, src, vs.Globals)
	return err
}
