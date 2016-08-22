package dbretrieval

import (
	"database/sql"
)

type FeatureResult struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	SieveableQuery string `json:"sieveable_query"`
}

func getFeatures(db *sql.DB, query string, packageName string) ([]FeatureResult, error) {
	rows, err := db.Query(query, packageName)
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

// Given a package name, return all app's features throughout its history
func GetFeaturesByPackageName(db *sql.DB, packageName string) ([]FeatureResult, error) {
	var query string = "SELECT name, description, sieveable_query FROM feature " +
		"WHERE id IN (SELECT feature_id FROM app_feature WHERE app_id IN " +
		"(SELECT id FROM app WHERE package_name = ?))"
	return getFeatures(db, query, packageName)
}

// Given a package name, return the latest app's features
func GetLatestFeaturesByPackageName(db *sql.DB, packageName string) ([]FeatureResult, error) {
	var query string = "SELECT name, description, sieveable_query FROM feature " +
		"WHERE id IN (SELECT feature_id FROM app_feature WHERE app_id IN " +
		"(SELECT id FROM app WHERE release_date =( " +
		"SELECT MAX(release_date) FROM app WHERE package_name=?)))"
	return getFeatures(db, query, packageName)
}

// Given a category name, return all features within that category
func GetFeaturesByCategoryName(db *sql.DB, categoryName string) ([]FeatureResult, error) {
	rows, err := db.Query("SELECT name, description, sieveable_query "+
		"FROM feature WHERE category_id IN "+
		"(SELECT id FROM category WHERE name = ?)", categoryName)
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

// Given a feature's name, return its details
func GetFeatureByName(db *sql.DB, featureName string) (FeatureResult, error) {
	var f FeatureResult
	err := db.QueryRow("SELECT name, description, sieveable_query "+
		"FROM feature WHERE name = ?", featureName).
		Scan(&f.Name, &f.Description, &f.SieveableQuery)
	if err != nil {
		return f, err
	}
	return f, nil
}
