package main

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(db *sql.DB) *httprouter.Router {
	router := httprouter.New()
	dbHandler := &DbHandler{db}
	router.GET("/apps/features/:featureName", dbHandler.getAppsByFeatureName)
	router.GET("/features/apps/:packageName", dbHandler.getAppFeaturesByPackageName)
	router.GET("/features/categories/:categoryName", dbHandler.getFeaturesByCategoryName)
	return router
}
