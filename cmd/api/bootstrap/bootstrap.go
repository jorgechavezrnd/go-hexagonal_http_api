package bootstrap

import "github.com/jorgechavezrnd/go-hexagonal_http_api/internal/platform/server"

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}
