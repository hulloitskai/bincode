package handler

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cockroachdb/errors"
	echo "github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"go.stevenxie.me/api/pkg/httputil"
	"go.stevenxie.me/bincode/encode"
)

// EncodeHandler returns an echo.HandlerFunc handles requests to encode an
// incoming file to a binary representation format.
//
// It accepts a form with the following fields:
//   key  – An alphabetical key to encode the data with.
//   text – Text to encode, mutually exclusive with the 'file' field.
//   file – File to encode, mutually exclusive with the 'text' field.
func EncodeHandler(log logrus.FieldLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse form fields.
		fields, err := c.FormParams()
		if err != nil {
			return errors.Wrap(err, "extracting form parameters")
		}

		var key string
		{
			const field = "key"
			if key = c.FormValue(field); key == "" {
				httputil.SetEchoStatusCode(c, http.StatusBadRequest)
				return errors.Newf("missing form field '%s'", field)
			}
		}

		var src io.ReadCloser
		{
			const (
				textField = "text"
				fileField = "file"
			)
			var (
				_, hasTextField = fields[textField]
				_, hasFileField = fields[fileField]
			)
			if hasTextField && hasFileField {
				return errors.Newf(
					"only one of the fields '%s' and '%s' can be present",
					fileField, textField,
				)
			}

			// Derive src from text / file field.
			if hasTextField {
				text := c.FormValue(textField)
				src = ioutil.NopCloser(strings.NewReader(text))
			} else {
				header, err := c.FormFile(fileField)
				if err != nil {
					if errors.Is(err, http.ErrMissingFile) {
						httputil.SetEchoStatusCode(c, http.StatusBadRequest)
						return errors.Newf(
							"one of form fields '%s' or '%s' must be present",
							fileField, textField,
						)
					}
					log.WithError(err).Error("Failed to read from file.")
					return errors.Wrap(err, "reading form file")
				}

				if src, err = header.Open(); err != nil {
					log.WithError(err).Error("Failed to open file.")
					return errors.Wrap(err, "opening form file")
				}
				defer src.Close()
			}
		}

		// Encode file into buf.
		var buf bytes.Buffer
		{
			enc, err := encode.NewEncoder(key, &buf)
			if err != nil {
				log.WithError(err).Error("Creating encoder.")
				return errors.Wrap(err, "creating encoder")
			}
			if _, err = io.Copy(enc, src); err != nil {
				log.WithError(err).Error("Error while encoding file.")
				return err
			}
		}

		// Close file.
		if err := src.Close(); err != nil {
			log.WithError(err).Error("Failed to close form file.")
			return errors.Wrap(err, "closing form file")
		}

		// Respond with buffer contents as plain text.
		return c.Blob(http.StatusOK, echo.MIMETextPlainCharsetUTF8, buf.Bytes())
	}
}
