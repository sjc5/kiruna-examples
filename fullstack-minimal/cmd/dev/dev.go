package main

import (
	"github.com/sjc5/kiruna"
	root "github.com/sjc5/kiruna-examples/fullstack-minimal"
)

func main() {
	devConfig := &kiruna.DevConfig{
		HealthcheckURL:    "http://localhost:8080/healthz",
		RefreshServerPort: 8081,
		WatchedFiles:      kiruna.WatchedFiles{".go.html": {}},
	}

	root.Kiruna.MustStartDev(devConfig)
}
