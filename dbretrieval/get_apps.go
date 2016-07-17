package dbretrieval

import (
	"database/sql"
	"time"
)

type AppResult struct {
	Id          string    `json:"id"`
	PackageName string    `json:"package_name"`
	VersionCode int       `json:"version_code"`
	VersionName string    `json:"version_name"`
	Downloads   int       `json:"downloads"`
	Ratings     float32   `json:"ratings"`
	ReleaseDate time.Time `json:"release_date"`
}

// Given a feature name, return all apps that have the given feature
func GetAppsByFeatureName(db *sql.DB, featureName string) ([]AppResult, error) {
	rows, err := db.Query("SELECT * FROM app WHERE id IN "+
		"(SELECT app_id FROM app_feature WHERE feature_id IN "+
		"(SELECT id FROM feature WHERE name = ? ))", featureName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	apps := []AppResult{}
	var timeLayout string = "2006-01-02"
	for rows.Next() {
		var app AppResult = AppResult{}
		var strDate string
		err := rows.Scan(&app.Id, &app.PackageName, &app.VersionCode,
			&app.VersionName, &app.Downloads, &app.Ratings, &strDate)
		if err != nil {
			return nil, err
		}
		t, err := time.Parse(timeLayout, strDate)
		if err != nil {
			return nil, err
		}
		app.ReleaseDate = t
		apps = append(apps, app)
	}
	return apps, nil
}
