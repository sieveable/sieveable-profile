package dbwriter

import (
	"encoding/json"
	"fmt"
	"time"
)

/* Example
{
    "category": {
        "type": "ui",
        "name": "material-design",
        "description": "Apps that implement material design components"
    },
    "feature": {
       "name": "feature_name",
       "description": "Feature description",
       "sieveable_query": "MATCH app WHERE <Button/> RETURN app"
    },
    "apps": [
	       {
			   "id":"com.example.app-8",
			   "packageName":"com.example.app",
			   "versionName":"1.2",
			   "versionCode":8,
			   "listing":{
				   "releaseDate":"2016-01-15",
				   "downloads":100
				   "ratings": 4.2
			   },
			   "ui":{},
			   "manifest":{}
		   }
	]
}

*/
type CustomTime struct {
	time.Time
}

const timeLayout = "2006-01-02"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	ct.Time, err = time.Parse(timeLayout, string(b))
	return
}

type ListingType struct {
	Downloads   int
	Ratings     float32
	ReleaseDate CustomTime
}
type AppType struct {
	Id          string
	PackageName string
	VersionCode int
	VersionName string
	Listing     ListingType
}
type CategoryType struct {
	Type        string
	Name        string
	Description string
}
type FeatureType struct {
	Name           string
	Description    string
	SieveableQuery string `json:"sieveable_query"`
}

type Response struct {
	Category CategoryType
	Feature  FeatureType
	Apps     []AppType
}

func Parse(fileContent *[]byte) (Response, error) {
	var res Response
	if err := json.Unmarshal(*fileContent, &res); err != nil {
		return res, fmt.Errorf("Failed to parse JSON data. %v\n", err)
	}
	return res, nil
}
