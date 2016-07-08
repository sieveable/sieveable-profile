package dbwriter

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var category = CategoryType{Type: "ui", Name: "material-design",
	Description: "Apps that implement Material Design"}
var feature = FeatureType{Name: "feature_name",
	Description: "feature_description", SieveableQuery: "sieveable_query_value"}
var releaseDate, _ = time.Parse("January 2, 2006", "January 16, 2016")
var cDate = CustomTime{releaseDate}
var listing = ListingType{Downloads: 100, Ratings: 4.2,
	ReleaseDate: cDate}
var app = AppType{Id: "com.example.app-8", PackageName: "com.example.app",
	VersionName: "1.2", VersionCode: 8, Listing: listing}

func TestInsertCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error occurred: '%s' when attempting to open a stub db connection", err)
	}
	defer db.Close()
	mock.ExpectPrepare("INSERT INTO category").ExpectExec().
		WithArgs(category.Name, category.Type, category.Description).
		WillReturnResult(sqlmock.NewResult(1, 1))
	categoryId, err := insertCategory(db, &category)
	if err != nil {
		t.Errorf("error was not expected while inserting a category: %s", err)
	}
	if categoryId != 1 {
		t.Error("Expected to return a category Id equals to 1, but got", categoryId)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}
func TestInsertFeature(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error occurred: '%s' when attempting to open a stub db connection", err)
	}
	defer db.Close()
	mock.ExpectPrepare("INSERT INTO feature").ExpectExec().
		WithArgs(feature.Name, feature.Description,
			feature.SieveableQuery, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	var categoryId int64 = 1
	featureId, err := insertFeature(db, &feature, &categoryId)
	if err != nil {
		t.Errorf("error was not expected while inserting a feature: %s", err)
	}
	if featureId != 1 {
		t.Error("Expected to return a feature Id equals to 1, but got", featureId)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}
func TestInsertApp(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error occurred: '%s' when attempting to open a stub db connection", err)
	}
	defer db.Close()
	mock.ExpectPrepare("INSERT INTO app").ExpectExec().
		WithArgs(app.Id, app.PackageName, app.VersionCode,
			app.VersionName, app.Listing.Downloads, app.Listing.Ratings,
			app.Listing.ReleaseDate.Format("2006-01-02")).
		WillReturnResult(sqlmock.NewResult(1, 1))
	if err = insertApp(db, &app); err != nil {
		t.Errorf("error was not expected while inserting an app: %s", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}
func TestInsertAppFeature(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error occurred: '%s' when attempting to open a stub db connection", err)
	}
	defer db.Close()
	mock.ExpectPrepare("INSERT INTO app_feature").ExpectExec().
		WithArgs("com.app", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	if err := insertAppFeature(db, "com.app", 1); err != nil {
		t.Errorf("error was not expected while inserting an app feature: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}
func TestIntegration(t *testing.T) {
	var res Response = Response{category, feature, []AppType{app}}
	if err := Insert(res); err != nil {
		t.Errorf("Expected to insert app feature but got an error instead: %v", err)
	}
}
