package main

import (
	"testing"

	"github.com/Schattenbrot/bookings/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Error("type is not *chi.Mux")
	}
}
