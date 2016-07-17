package dbretrieval

import (
	"database/sql"
)

type FeatureResult struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	SieveableQuery string `json:"sieveable_query"`
}

// Given the package name, return the app's features
func GetFeaturesByPackageName(db *sql.DB, packageName string) ([]FeatureResult, error) {
	rows, err := db.Query("SELECT name, description, sieveable_query FROM feature "+
		"WHERE id IN (SELECT feature_id FROM app_feature WHERE app_id IN "+
		"(SELECT id FROM app WHERE package_name = ?))", packageName)
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
