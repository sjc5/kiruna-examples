package main

import "github.com/sjc5/kiruna"

func main() {
	k := &kiruna.Kiruna{
		Config: &kiruna.Config{
			EntryPoint: "cmd/app/main.go",
			DevConfig: &kiruna.DevConfig{
				ServerOnly: true,
			},
		},
	}
	k.Dev()
}
