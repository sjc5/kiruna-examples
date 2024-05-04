package root

import "github.com/sjc5/kiruna"

var Kiruna *kiruna.Kiruna

func init() {
	Kiruna = &kiruna.Kiruna{
		Config: &kiruna.Config{
			EntryPoint: "cmd/app/main.go",
		},
	}
}
