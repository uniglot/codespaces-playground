package main

import (
	"net/http"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			NewHTTPServer,
			NewServeMux,
			NewEchoHandler,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
