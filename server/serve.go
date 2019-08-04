package server

import (
	"github.com/cockroachdb/errors"
)

// ListenAndServe listens and serves on an API for encoding and decoding binary
// representation formats.
func (srv *Server) ListenAndServe(addr string) error {
	if err := srv.registerRoutes(); err != nil {
		return errors.Wrap(err, "server: registering routes")
	}

	srv.log.WithField("addr", addr).Info("Listening for requests...")
	return srv.echo.Start(addr)
}
