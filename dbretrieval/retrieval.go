package dbretrieval

import (
	"database/sql"
	"fmt"
)

type FeatureResult struct {
	Name           string
	Description    string
	SieveableQuery string
}

// Given the package name, return the app's features
func GetAppFeaturesByPackageName(db *sql.DB, packageName string) ([]FeatureResult, error) {
	fmt.Printf("\t\t%s\n", packageName)
	rows, err := db.Query("SELECT name, description, sieveable_query FROM feature "+
		"WHERE id IN (SELECT feature_id FROM app_feature WHERE app_id IN "+
		"(SELECT id FROM app WHERE package_name=?))", packageName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	features := []FeatureResult{}
	for rows.Next() {
		var f FeatureResult = FeatureResult{}
		err := rows.Scan(&f.Name, &f.Description, &f.SieveableQuery)
		if err != nil {
			return nil, err
		}
		features = append(features, f)
	}
	return features, nil
}

/*
// Given a feature name, return all apps that have the given feature
func GetAppsWithFeature(db *sql.DB, featureName string) ([]AppType, error){
}

// Given a category, return all features within that category
func GetFeaturesByCategory(db *sql.DB, categoryName string) ([]FeatureType, error) {
}
// Given a cateogry, return all apps that implement features within that category
func GetAppsByFeatureCategory(db *sql.DB, categoryName string)([]AppType, error) {

}
*/
