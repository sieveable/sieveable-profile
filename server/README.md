# Sieveable Profile API

## Apps
### List apps by a feature name
`GET /apps/features/:featureName`
### List top apps by a feature name
`GET /apps/features/:featureName?top=true`

## Features
### List feature details by name
`GET /feature/:featureName`
### List all features by an app's package name
`GET /features/apps/:packageName`
### List features by an app's package name(only limit to the latest app version)
`GET /features/apps/:packageName?latest=true`

### List features by their category name
`GET /features/categories/:categoryName`

