package dbretrieval

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sieveable/sieveable-profile/dbwriter"
	"os"
	"testing"
	"time"
)

var db *sql.DB
var category = dbwriter.CategoryType{Type: "ui", Name: "material-design",
	Description: "Apps that implement Material Design"}
var firstFeature = dbwriter.FeatureType{Name: "first_feature_name",
	Description: "feature_description", SieveableQuery: "sieveable_query_value"}
var secondFeature = dbwriter.FeatureType{Name: "second_feature_name",
	Description: "feature_description", SieveableQuery: "sieveable_query_value"}
var firstReleaseDate, _ = time.Parse("January 2, 2006", "January 16, 2016")
var secondReleaseDate, _ = time.Parse("January 2, 2006", "March 03 , 2016")
var firstListing = dbwriter.ListingType{Downloads: 100, Ratings: 4.2,
	ReleaseDate: dbwriter.CustomTime{firstReleaseDate}}
var secondListing = dbwriter.ListingType{Downloads: 100, Ratings: 4.2,
	ReleaseDate: dbwriter.CustomTime{secondReleaseDate}}
var firstApp = dbwriter.AppType{Id: "com.example.app-8", PackageName: "com.example.app",
	VersionName: "1.2", VersionCode: 8, Listing: firstListing}
var secondApp = dbwriter.AppType{Id: "com.example.app-9", PackageName: "com.example.app",
	VersionName: "1.3", VersionCode: 9, Listing: secondListing}

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		fmt.Printf("Failed to setup test. %v", err)
		os.Exit(1)
	}
	defer db.Close()
	os.Exit(m.Run())
}

func setup() (err error) {
	db, err = getDbConnection()
	if err != nil {
		return err
	}
	var firstRes dbwriter.Response = dbwriter.Response{category,
		firstFeature, []dbwriter.AppType{firstApp}}
	var secondRes dbwriter.Response = dbwriter.Response{category,
		secondFeature, []dbwriter.AppType{secondApp}}
	if err := dbwriter.Insert(db, firstRes); err != nil {
		return err
	}
	return dbwriter.Insert(db, secondRes)
}

func getDbConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("USER")+":"+os.Getenv("PW")+"@/"+os.Getenv("DB"))
	if err != nil {
		return nil, fmt.Errorf("Failed to get a handle for the database %s. %v\n",
			os.Getenv("DB"), err.Error())
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to connect to DB. Make sure that the required "+
			"environment variables are set. %v\n", err.Error())
	}
	return db, nil
}

func TestGetFeaturesByPackageName(t *testing.T) {
	features, err := GetFeaturesByPackageName(db, firstApp.PackageName)
	if err != nil {
		t.Errorf("Expected app features but got an error instead. %v", err)
	}
	if len(features) != 2 {
		t.Errorf("Expected an array of features with size 2 but got %d instead", len(features))
	}
}

func TestGetLatestFeaturesByPackageName(t *testing.T) {
	features, err := GetLatestFeaturesByPackageName(db, firstApp.PackageName)
	if err != nil {
		t.Errorf("Expected app features but got an error instead. %v", err)
	}
	if len(features) != 1 {
		t.Errorf("Expected an array of features with size 1 but got %d instead", len(features))
	}
}

func TestGetAppsByFeatureName(t *testing.T) {
	apps, err := GetAppsByFeatureName(db, firstFeature.Name)
	if err != nil {
		t.Errorf("Expected an array of apps but got an error instead. %v", err)
	}
	if len(apps) != 1 {
		t.Errorf("Expected an array of apps with size 1 but got %d instead", len(apps))
	}
}

func TestGetFeaturesByCategoryName(t *testing.T) {
	features, err := GetFeaturesByCategoryName(db, category.Name)
	if err != nil {
		t.Errorf("Expected an array of features but got an error instead. %v", err)
	}
	if len(features) != 2 {
		t.Errorf("Expected an array of features with size 2 but got %d instead", len(features))
	}
}
