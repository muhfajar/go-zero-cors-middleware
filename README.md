# Go-Zero Middleware to handle CORS request
[![Build Status](https://app.travis-ci.com/muhfajar/go-zero-cors-middleware.svg?branch=main&status=passed)](https://app.travis-ci.com/muhfajar/go-zero-cors-middleware)
[![codecov](https://codecov.io/gh/muhfajar/go-zero-cors-middleware/branch/main/graph/badge.svg?token=1FB5E5PDCH)](https://codecov.io/gh/muhfajar/go-zero-cors-middleware)

## Getting Started

Install `middleware`:

    go get github.com/muhfajar/go-zero-cors-middleware


After setting up your [Go-Zero](https://go-zero.dev/en/) project, update your `main` package to register CORS middleware.

```
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
	
	// Register go-zero-cors-middleware handler to handle preflight request
	cors := middleware.NewCORSMiddleware(&middleware.Options{})
	
	// Add run option WithNotAllowedHandler and register `.Handler()` to handle `OPTIONS` request (preflight)
	server := rest.MustNewServer(c.RestConf,
	    rest.WithNotAllowedHandler(cors.Handler()),
	)

	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// Register go-zero-cors-middleware
	server.Use(cors.Handle)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
```

Available options

```
AllowCredentials bool       Indicates whether the request can include user credentials like cookies, HTTP authentication or client side SSL certificates.
AllowHeaders     []string   A list of non simple headers the client is allowed to use with cross-domain requests.
AllowMethods     []string   A list of methods the client is allowed to use with cross-domain requests.
ExposeHeaders    []string   Indicates which headers are safe to expose to the API of a CORS API specification.
```

Default value

```
AllowCredentials    false
AllowHeaders        []string{"Content-Type", "X-CSRF-Token", "Authorization", "AccessToken", "Token"}
AllowMethods        []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}
ExposeHeaders       []string{"Content-Length", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"}
```

By default, if request `Origin` header value is null or not included in request and `AllowCredentials` is not set in options, `Access-Control-Allow-Origin` will be return `*`.
But if `AllowCredentials` set by `true` value and request `Origin` header value is present, `Access-Control-Allow-Origin` will be reflected by the request `Origin` value.

## Licenses

All source code is licensed under the [MIT License](https://github.com/muhfajar/go-zero-cors-middleware/main/LICENSE).