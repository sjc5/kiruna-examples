package global

import (
	"embed"

	"github.com/sjc5/kiruna"
)

//go:embed dist/kiruna
var DistFS embed.FS

var Kiruna *kiruna.Kiruna

func init() {
	Kiruna = &kiruna.Kiruna{
		Config: &kiruna.Config{
			DistFS:     DistFS,
			EntryPoint: "cmd/app/main.go",
			DevConfig: &kiruna.DevConfig{
				HealthcheckURL:    "http://localhost:8080/healthz",
				RefreshServerPort: 8081,
				WatchedFiles:      kiruna.WatchedFiles{".go.html": {}},
			},
		},
	}
}
