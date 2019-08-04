package server

import (
	"github.com/sirupsen/logrus"

	"go.stevenxie.me/bincode/internal/info"
	serverinfo "go.stevenxie.me/bincode/server/internal/info"

	"go.stevenxie.me/api/pkg/httputil"
	"go.stevenxie.me/bincode/server/handler"
)

func (srv *Server) registerRoutes() error {
	e := srv.echo
	e.HTTPErrorHandler = httputil.ErrorHandler(srv.hlog("error"))

	// Register routes.
	e.GET("/", httputil.InfoHandler(serverinfo.Name, info.Version))
	e.POST("/encode", handler.EncodeHandler(srv.hlog("encode")))
	e.POST("/decode", handler.DecodeHandler(srv.hlog("decode")))

	return nil
}

func (srv *Server) hlog(name string) logrus.FieldLogger {
	return srv.log.WithField("handler", name)
}
