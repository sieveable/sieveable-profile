package main

import (
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/sieveable/sieveable-profile/dbretrieval"
	"net/http"
)

type DbHandler struct {
	db *sql.DB
}

func (dbHandler *DbHandler) getAppsByFeatureName(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	apps, err := dbretrieval.GetAppsByFeatureName(dbHandler.db, ps.ByName("featureName"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apps); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (dbHandler *DbHandler) getAppFeaturesByPackageName(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var features []dbretrieval.FeatureResult
	var err error
	queryValues := r.URL.Query()
	if queryValues.Get("latest") == "true" {
		features, err = dbretrieval.GetLatestFeaturesByPackageName(dbHandler.db,
			ps.ByName("packageName"))
	} else {
		features, err = dbretrieval.GetFeaturesByPackageName(dbHandler.db,
			ps.ByName("packageName"))
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(features); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (dbHandler *DbHandler) getFeaturesByCategoryName(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	features, err := dbretrieval.GetFeaturesByCategoryName(dbHandler.db,
		ps.ByName("categoryName"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(features); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
