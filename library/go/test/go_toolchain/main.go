package main

import (
	"github.com/ydb-platform/nbs/library/go/test/go_toolchain/gotoolchain"
	"github.com/ydb-platform/nbs/library/go/test/recipe"
)

type goToolchain struct{}

func (r *goToolchain) Start() error {
	setEnv := func(k, v string) error {
		recipe.SetEnv(k, v)
		return nil
	}

	return gotoolchain.Setup(setEnv)
}

func (r *goToolchain) Stop() error {
	return nil
}

func main() {
	recipe.Run(&goToolchain{})
}
