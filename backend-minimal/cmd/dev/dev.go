package main

import (
	"github.com/sjc5/kiruna"
	root "github.com/sjc5/kiruna-examples/backend-minimal"
)

func main() {
	root.Kiruna.MustStartDev(&kiruna.DevConfig{
		ServerOnly: true,
	})
}
