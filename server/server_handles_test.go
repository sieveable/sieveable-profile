package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/sieveable/sieveable-profile/dbretrieval"
	"github.com/sieveable/sieveable-profile/dbwriter"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var db *sql.DB
var category = dbwriter.CategoryType{Type: "ui", Name: "material-design",
	Description: "Apps that implement Material Design"}
var feature = dbwriter.FeatureType{Name: "feature_name",
	Description: "feature_description", SieveableQuery: "sieveable_query_value"}
var releaseDate, _ = time.Parse("January 2, 2006", "January 16, 2016")
var cDate = dbwriter.CustomTime{releaseDate}
var listing = dbwriter.ListingType{Downloads: 100, Ratings: 4.2,
	ReleaseDate: cDate}
var app = dbwriter.AppType{Id: "com.example.app-8", PackageName: "com.example.app",
	VersionName: "1.2", VersionCode: 8, Listing: listing}

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		fmt.Printf("Failed to setup test. %v", err)
		os.Exit(1)
	}
	defer db.Close()
	os.Exit(m.Run())
}

func setup() (err error) {
	db, err = getDatabaseConnection()
	if err != nil {
		return err
	}
	var res dbwriter.Response = dbwriter.Response{category, feature, []dbwriter.AppType{app}}
	return dbwriter.Insert(db, res)
}

func getDatabaseConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("USER")+":"+os.Getenv("PW")+"@/"+os.Getenv("DB"))
	if err != nil {
		return nil, fmt.Errorf("Failed to get a handle for the database. %v\n", err.Error())
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to connect to DB. Make sure that the required "+
			"environment variables are set. %v\n", err.Error())
	}
	return db, nil
}

func doHttpRequest(uri string, ps httprouter.Params, handle httprouter.Handle) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	res := httptest.NewRecorder()
	handle(res, req, ps)
	return res, nil
}

func TestGetAppsByFeatureName(t *testing.T) {
	dbHandler := &DbHandler{db}
	ps := httprouter.Params{httprouter.Param{"featureName", "feature_name"}}
	uri := "/apps/features/"
	resp, err := doHttpRequest(uri, ps, dbHandler.getAppsByFeatureName)
	if err != nil {
		t.Error(err)
	}
	if resp.Code != 200 {
		t.Errorf("Expected HTTP response code of 200 but got %d", resp.Code)
	}
	var results []dbretrieval.AppResult = []dbretrieval.AppResult{}
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		t.Errorf("Failed to decode response body. %v", err)
	}
	if len(results) != 1 {
		t.Errorf("Expected the length of the result array to be one but got %d", len(results))
	}
	if results[0].PackageName != "com.example.app" {
		t.Errorf("Expected app's package name to equal com.example.app but got %s",
			results[0].PackageName)
	}
}

func TestGetAppFeaturesByPackageName(t *testing.T) {
	dbHandler := &DbHandler{db}
	ps := httprouter.Params{httprouter.Param{"packageName", "com.example.app"}}
	uri := "/features/apps/"
	resp, err := doHttpRequest(uri, ps, dbHandler.getAppFeaturesByPackageName)
	if err != nil {
		t.Error(err)
	}
	if resp.Code != 200 {
		t.Errorf("Expected HTTP response code of 200 but got %d", resp.Code)
	}
	var results []dbretrieval.FeatureResult = []dbretrieval.FeatureResult{}
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		t.Errorf("Failed to decode response body. %v", err)
	}
	if len(results) != 1 {
		t.Errorf("Expected the length of the result array to be one but got %d", len(results))
	}
	if results[0].Name != "feature_name" {
		t.Errorf("Expected feature name to equal feature_name but got %s", results[0].Name)
	}
}

func TestGetLatestAppFeaturesByPackageName(t *testing.T) {
	dbHandler := &DbHandler{db}
	ps := httprouter.Params{httprouter.Param{"packageName", "com.example.app"}}
	uri := "/features/apps?latest=true"
	resp, err := doHttpRequest(uri, ps, dbHandler.getAppFeaturesByPackageName)
	if err != nil {
		t.Error(err)
	}
	if resp.Code != 200 {
		t.Errorf("Expected HTTP response code of 200 but got %d", resp.Code)
	}
	var results []dbretrieval.FeatureResult = []dbretrieval.FeatureResult{}
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		t.Errorf("Failed to decode response body. %v", err)
	}
	if len(results) != 1 {
		t.Errorf("Expected the length of the result array to be one but got %d", len(results))
	}
	if results[0].Name != "feature_name" {
		t.Errorf("Expected feature name to equal feature_name but got %s", results[0].Name)
	}
}

func TestGetAppFeaturesByCategoryName(t *testing.T) {
	dbHandler := &DbHandler{db}
	ps := httprouter.Params{httprouter.Param{"categoryName", "material-design"}}
	uri := "/features/apps/"
	resp, err := doHttpRequest(uri, ps, dbHandler.getFeaturesByCategoryName)
	if err != nil {
		t.Error(err)
	}
	if resp.Code != 200 {
		t.Errorf("Expected HTTP response code of 200 but got %d", resp.Code)
	}
	var results []dbretrieval.FeatureResult = []dbretrieval.FeatureResult{}
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		t.Errorf("Failed to decode response body. %v", err)
	}
	if len(results) != 1 {
		t.Errorf("Expected the length of the result array to be one but got %d", len(results))
	}
	if results[0].Name != "feature_name" {
		t.Errorf("Expected feature name to equal feature_name but got %s", results[0].Name)
	}
}
