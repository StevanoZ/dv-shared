package shrd_middleware

import (
	"time"

	shrd_utils "github.com/StevanoZ/dv-shared/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupMiddleware(route *chi.Mux, config *shrd_utils.BaseConfig) {
	if config.Environment == shrd_utils.LOCAL || config.Environment == shrd_utils.TEST {
		route.Use(Recovery)
	} else {
		route.Use(middleware.RequestID)
		route.Use(middleware.RealIP)
		route.Use(middleware.Logger)
		route.Use(middleware.Timeout(60 * time.Second))

		route.Use(Recovery)
	}
}
