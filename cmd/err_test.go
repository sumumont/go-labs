package main

import (
	"errors"
	"github.com/go-labs/internal/logging"
	"github.com/rs/zerolog/hlog"
	"net/http"
	"testing"
)

func TestErr(t *testing.T) {
	err()
}

func err() {
	hlog.RequestIDHandler("uuid", "uuid")(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := hlog.FromRequest(r)
			log.Info().Msg("hello from TestRequestID")
		}),
	)

	logging.Error(errors.New("dsadasdsad")).Send()
	logging.Debug().Msg("1")
	logging.Debug().Msg("2")
	logging.Debug().Msg("3")
}
