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
	db, err = getDbConnection()
	if err != nil {
		return err
	}
	var res dbwriter.Response = dbwriter.Response{category, feature, []dbwriter.AppType{app}}
	return dbwriter.Insert(db, res)
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

func TestGetAppFeatures(t *testing.T) {
	features, err := GetAppFeaturesByPackageName(db, app.PackageName)
	if err != nil {
		t.Errorf("Expected app features but got an error instead. %v", err)
	}
	if len(features) != 1 {
		t.Errorf("Expected an array of features with size 1 but got %d instead", len(features))
	}
}
