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
	queryValues := r.URL.Query()
	if queryValues.Get("top") == "true" {
		topApps, err := dbretrieval.GetTopAppsByFeatureName(dbHandler.db, ps.ByName("featureName"))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(topApps); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		allApps, err := dbretrieval.GetAppsByFeatureName(dbHandler.db, ps.ByName("featureName"))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(allApps); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
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

func (dbHandler *DbHandler) getFeature(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	feature, err := dbretrieval.GetFeatureByName(dbHandler.db,
		ps.ByName("featureName"))
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Not found", 404)
		}
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(feature); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (dbHandler *DbHandler) getCategoriesByType(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	categories, err := dbretrieval.GetCategoriesByType(dbHandler.db,
		ps.ByName("type"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
