package routes

import (
	"database/sql"

	"github.com/gasBlar/GoGoManager/api/v1/middleware"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.JWTMiddleware)
	r.Use(middleware.LogMiddleware)
	middleware.PromotheusInit()
	s := r.PathPrefix("/api/v1").Subrouter()

	r.Handle("/metrics", promhttp.Handler())
	r.Use(middleware.TrackMetrics)

	ExampleRoutes(s)
	AuthRoutes(s)
	UserRoutes(s)
	EmployeeRoutes(db, s)
	DepartmentRoutes(db, s)

	return r
}
