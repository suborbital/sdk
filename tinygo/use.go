package tinygo

import (
	"github.com/suborbital/sdk/tinygo/internal/ffi"
	"github.com/suborbital/sdk/tinygo/plugin"
)

func Use(plugin plugin.Plugin) {
	ffi.Use(plugin)
}
