package dbwriter

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	var jsonBlob = []byte(`{
		"category": {
			"type": "ui",
			"name": "material-design",
			"description": "Apps that implement Material Design"
		},
		"feature": {
			"name": "feature_name",
			"description": "feature_description",
			"sieveable_query": "sieveable_query_value"
		},
		"apps": [
		{
			"id":"com.example.app-8",
			"packageName":"com.example.app",
			"versionName":"1.2",
			"versionCode":8,
			"listing":{
				"releaseDate":"January 16, 2016",
				"downloads":100,
				"ratings": 4.2
			},
			"ui":{},
			"manifest":{}
		}
		]
	}`)
	parsed, err := Parse(&jsonBlob)
	if err != nil {
		t.Errorf("Expected to parse the JSON blob but got an error instead: %v", err)
	}
	expectedCategory := CategoryType{Type: "ui", Name: "material-design",
		Description: "Apps that implement Material Design"}
	expectedFeature := FeatureType{Name: "feature_name",
		Description: "feature_description", SieveableQuery: "sieveable_query_value"}
	releaseDate, _ := time.Parse("January 2, 2006", "January 16, 2016")
	cDate := CustomTime{releaseDate}
	expectedListing := ListingType{Downloads: 100, Ratings: 4.2,
		ReleaseDate: cDate}
	expectedApp := AppType{Id: "com.example.app-8", PackageName: "com.example.app",
		VersionName: "1.2", VersionCode: 8, Listing: expectedListing}
	if parsed.Category != expectedCategory {
		t.Errorf("\nExpected: %v\nActual: %v\n", expectedCategory, parsed.Category)
	}
	if parsed.Feature != expectedFeature {
		t.Errorf("\nExpected: %v\nActual: %v\n", expectedFeature, parsed.Feature)
	}
	if len(parsed.Apps) != 1 {
		t.Errorf("\nExpected: %d\nActual: %d\n", 1, len(parsed.Apps))
	}
	if parsed.Apps[0] != expectedApp {
		t.Errorf("\nExpected: %v\nActual: %v\n", expectedApp, parsed.Apps[0])
	}
}
