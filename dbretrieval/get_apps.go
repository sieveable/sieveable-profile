package dbretrieval

import (
	"database/sql"
	"github.com/sieveable/sieveable-profile/dbwriter"
	"time"
)

// Given a feature name, return all apps that have the given feature
func GetAppsByFeatureName(db *sql.DB, featureName string) ([]dbwriter.AppType, error) {
	rows, err := db.Query("SELECT * FROM app WHERE id IN "+
		"(SELECT app_id FROM app_feature WHERE feature_id IN "+
		"(SELECT id FROM feature WHERE name = ? ))", featureName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	apps := []dbwriter.AppType{}
	var timeLayout string = "2006-01-02"
	for rows.Next() {
		var app dbwriter.AppType = dbwriter.AppType{}
		var strDate string
		var customTime = dbwriter.CustomTime{}
		var listing = dbwriter.ListingType{}
		err := rows.Scan(&app.Id, &app.PackageName, &app.VersionCode,
			&app.VersionName, &listing.Downloads, &listing.Ratings, &strDate)
		if err != nil {
			return nil, err
		}
		customTime.Time, err = time.Parse(timeLayout, strDate)
		if err != nil {
			return nil, err
		}
		listing.ReleaseDate = customTime
		app.Listing = listing
		apps = append(apps, app)
	}
	return apps, nil
}
