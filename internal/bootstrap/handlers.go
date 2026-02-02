package bootstrap

import (
	transport_http "github.com/zmonsoonz/go-test-api/internal/transport/http"
)

type Handlers struct {
	Auth *transport_http.AuthHandler
	Track *transport_http.TrackHandler
}

func InitHandlers(services *Services) *Handlers {
	return &Handlers{
		Auth: transport_http.NewAuthHandler(services.Auth),
		Track: transport_http.NewTrackHandler(services.Track),
	}
}