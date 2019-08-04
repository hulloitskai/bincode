package server

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	"go.stevenxie.me/api/pkg/zero"
)

// New creates a new Server.
func New(opts ...func(*Config)) *Server {
	cfg := Config{
		Logger: zero.Logger(),
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	// Create Echo.
	echo := echo.New()
	echo.Logger.SetOutput(ioutil.Discard) // disable Echo logger

	// Configure middleware.
	echo.Pre(middleware.RemoveTrailingSlashWithConfig(
		middleware.TrailingSlashConfig{
			RedirectCode: http.StatusPermanentRedirect,
		},
	))

	// Enable Access-Control-Allow-Origin: * during development.
	if os.Getenv("GOENV") == "development" {
		echo.Use(middleware.CORS())
	}

	// Create server.
	return &Server{
		echo: echo,
		log:  cfg.Logger,
	}
}

type (
	// A Server serves an API for encoding and decoding binary representation
	// formats.
	Server struct {
		echo *echo.Echo
		log  logrus.FieldLogger

		enc io.Writer
		dec io.Reader
	}

	// A Config configures a Server.
	Config struct {
		Logger logrus.FieldLogger
	}
)

// Shutdown gracefully shuts down the server.
func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.echo.Shutdown(ctx)
}
