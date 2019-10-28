package server

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"go.stevenxie.me/gopkg/logutil"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	"go.stevenxie.me/api/v2/pkg/basic"
)

// New creates a new Server.
func New(opts ...basic.Option) *Server {
	cfg := basic.BuildConfig(opts...)
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
		log:  logutil.WithComponent(cfg.Logger, (*Server)(nil)),
	}
}

// A Server serves an API for encoding and decoding binary representation
// formats.
type Server struct {
	echo *echo.Echo
	log  *logrus.Entry

	enc io.Writer
	dec io.Reader
}

// Shutdown gracefully shuts down the server.
func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.echo.Shutdown(ctx)
}
