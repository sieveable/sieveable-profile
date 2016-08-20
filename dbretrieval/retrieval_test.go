package dbretrieval

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"testing"
)

var db *sql.DB

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
	return nil
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
	features, err := GetFeaturesByPackageName(db, "com.example.app")
	if err != nil {
		t.Errorf("Expected app features but got an error instead. %v", err)
	}
	if len(features) != 2 {
		t.Errorf("Expected an array of features with size 2 but got %d instead", len(features))
	}
}

func TestGetLatestFeaturesByPackageName(t *testing.T) {
	features, err := GetLatestFeaturesByPackageName(db, "com.example.app")
	if err != nil {
		t.Errorf("Expected app features but got an error instead. %v", err)
	}
	if len(features) != 1 {
		t.Errorf("Expected an array of features with size 1 but got %d instead", len(features))
	}
}

func TestGetAppsByFeatureName(t *testing.T) {
	apps, err := GetAppsByFeatureName(db, "first_feature_name")
	if err != nil {
		t.Errorf("Expected an array of apps but got an error instead. %v", err)
	}
	if len(apps) != 1 {
		t.Errorf("Expected an array of apps with size 1 but got %d instead", len(apps))
	}
}

func TestGetTopAppsByFeatureName(t *testing.T) {
	apps, err := GetTopAppsByFeatureName(db, "first_feature_name")
	if err != nil {
		t.Errorf("Expected an array of apps but got an error instead. %v", err)
	}
	if len(apps) != 1 {
		t.Errorf("Expected an array of apps with size 1 but got %d instead", len(apps))
	}
}

func TestGetFeaturesByCategoryName(t *testing.T) {
	features, err := GetFeaturesByCategoryName(db, "material-design")
	if err != nil {
		t.Errorf("Expected an array of features but got an error instead. %v", err)
	}
	if len(features) != 2 {
		t.Errorf("Expected an array of features with size 2 but got %d instead", len(features))
	}
}

func TestGetCategoriesByType(t *testing.T) {
	categories, err := GetCategoriesByType(db, "ui")
	if err != nil {
		t.Errorf("Expected feature categories but got an error instead. %v", err)
	}
	if len(categories) != 1 {
		t.Errorf("Expected an array of categories with size 1 but got %d instead", len(categories))
	}
}
