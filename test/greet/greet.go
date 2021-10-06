package main

import (
	"flag"
	"fmt"

	"github.com/muhfajar/go-zero-cors-middleware/test/greet/internal/config"
	"github.com/muhfajar/go-zero-cors-middleware/test/greet/internal/handler"
	"github.com/muhfajar/go-zero-cors-middleware/test/greet/internal/svc"

	middleware "github.com/muhfajar/go-zero-cors-middleware"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/greet-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	// Register cors handler to handle preflight request
	cors := middleware.NewCORSMiddleware(&middleware.Options{})
	server := rest.MustNewServer(c.RestConf, rest.WithNotAllowedHandler(
		cors.Handler(),
	))

	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// Register cors middleware
	server.Use(cors.Handle)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
