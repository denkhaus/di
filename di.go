package di

import (
	"github.com/samber/do/v2"
)

var injector do.Injector

func Injector() do.Injector {
	if injector == nil {
		injector = do.New()
	}

	return injector
}

func Shutdown() error {
	if injector != nil {
		return injector.Shutdown()
	}

	return nil
}
