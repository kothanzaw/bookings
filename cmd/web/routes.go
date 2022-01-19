package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kothanzaw/bookings/internal/config"
	"github.com/kothanzaw/bookings/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler{

	mux :=chi.NewRouter()
	
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)


	mux.Get("/",handlers.Repo.Home)
	mux.Get("/about",handlers.Repo.About)
	mux.Get("/general-quarters",handlers.Repo.General)
	mux.Get("/majors-suit",handlers.Repo.Majors)
	mux.Get("/contact",handlers.Repo.Contact)
	mux.Get("/search-availability",handlers.Repo.Availability)
	mux.Post("/search-availability",handlers.Repo.PostAvailability)
	mux.Get("/make-reservation",handlers.Repo.MakeReservation)
	mux.Post("/search-availability-json",handlers.Repo.AvailabilityJSON)


	fileServer :=http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*",http.StripPrefix("/static",fileServer))

	return mux
}