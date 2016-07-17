package main

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(db *sql.DB) *httprouter.Router {
	router := httprouter.New()
	dbHandler := &DbHandler{db}
	router.GET("/apps/features/:featureName", dbHandler.GetAppsByFeatureName)
	router.GET("/features/apps/:packageName", dbHandler.GetAppFeaturesByPackageName)
	router.GET("/features/categories/:categoryName", dbHandler.GetFeaturesByCategoryName)
	return router
}
