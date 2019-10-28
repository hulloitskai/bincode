package handler

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cockroachdb/errors"
	echo "github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"go.stevenxie.me/bincode/encode"
	"go.stevenxie.me/gopkg/logutil"
)

// DecodeHandler returns an echo.HandlerFunc handles requests to decode from
// a binary representation format.
func DecodeHandler(log *logrus.Entry) echo.HandlerFunc {
	log = logutil.WithComponent(log, DecodeHandler)
	return func(c echo.Context) error {
		body := c.Request().Body
		defer body.Close()

		// Decode request body.
		result, err := ioutil.ReadAll(encode.NewDecoder(body))
		if err != nil {
			return err
		}
		if err = body.Close(); err != nil {
			log.WithError(err).Error("Failed to close request body.")
			return errors.Wrap(err, "closing request body")
		}

		var zero time.Time
		http.ServeContent(
			c.Response().Writer,
			c.Request(),
			"result",
			zero,
			bytes.NewReader(result),
		)
		return nil
	}
}
