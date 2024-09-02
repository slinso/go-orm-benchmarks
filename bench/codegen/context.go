package codegen

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Context struct {
	Log *zerolog.Logger
}

func NewContext() *Context {
	disabled := log.Logger.Level(zerolog.Disabled)

	return &Context{
		Log: &disabled,
	}
}
