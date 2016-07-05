package dbwriter

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strings"
)

// Insert the parsed response into MySQL db
func Insert(res Response) {
	db, err := sql.Open("mysql", os.Getenv("USER")+":"+os.Getenv("PW")+"@/"+os.Getenv("DB"))
	if err != nil {
		log.Fatalf("Failed to get a handle for the database %s. %v\n", os.Getenv("DB"), err.Error())
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to connect to DB. Make sure that the required "+
			"environment variables are set. %v\n", err.Error())
	}
	log.Println("Inserting new category", res.Category.Name)
	categoryId := insertCategory(db, &res.Category)
	log.Println("Inserting new feature", res.Feature.Name)
	featureId := insertFeature(db, &res.Feature, &categoryId)
	for _, app := range res.Apps {
		log.Printf("Inserting features for %s\n", app.Id)
		insertApp(db, &app)
		insertAppFeature(db, app.Id, featureId)
	}
}

func insertCategory(db *sql.DB, Category *CategoryType) (categoryId int64) {
	stmtIns, err := db.Prepare("INSERT INTO category SET name=?,type=?,description=?")
	checkError("category", err)
	defer stmtIns.Close()
	execResult, err := stmtIns.Exec(Category.Name, Category.Type, Category.Description)
	checkError("category", err)
	if execResult == nil {
		categoryId = getIdByName(db, "category", Category.Name)
	} else {
		categoryId, err = execResult.LastInsertId()
		if err != nil {
			log.Fatalln("Failed to get category id")
		}
	}
	return
}
func getIdByName(db *sql.DB, table string, name string) (id int64) {
	if table != "category" && table != "feature" {
		log.Fatalf("Cannot get id for an unknown table.")
	}
	stmtOut, err := db.Prepare("SELECT id FROM " + table + " WHERE name =?")
	if err != nil {
		log.Fatalf("Failed to find "+table+" id for %s\n", name)
	}
	defer stmtOut.Close()
	err = stmtOut.QueryRow(name).Scan(&id)
	if err != nil {
		log.Fatalf("Failed to find "+table+" id for %s\n", name)
	}
	return
}
func insertFeature(db *sql.DB, feature *FeatureType, categoryId *int64) (featureId int64) {
	stmtIns, err := db.Prepare("INSERT INTO feature SET " +
		"name=?,description=?,sieveable_query=?,category_id=?")
	checkError("feature", err)
	defer stmtIns.Close()
	execResult, err := stmtIns.Exec(feature.Name, feature.Description,
		feature.SieveableQuery, categoryId)
	checkError("feature", err)
	if execResult == nil {
		featureId = getIdByName(db, "feature", feature.Name)
	} else {
		featureId, err = execResult.LastInsertId()
		if err != nil {
			log.Fatalln("Failed to get feature id")
		}
	}
	return
}
func insertApp(db *sql.DB, app *AppType) {
	stmtIns, err := db.Prepare("INSERT INTO app SET id=?, package_name=?," +
		"version_code=?,version_name=?,downloads=?,ratings=?,release_date=?")
	checkError("app", err)
	defer stmtIns.Close()
	_, err = stmtIns.Exec(app.Id, app.PackageName, app.VersionCode,
		app.VersionName, app.Listing.Downloads, app.Listing.Ratings,
		app.Listing.ReleaseDate.Format("2006-01-02"))
	checkError("app", err)
}
func insertAppFeature(db *sql.DB, appId string, featureId int64) {
	stmtIns, err := db.Prepare("INSERT INTO app_feature VALUES(?,?)")
	checkError("app_feature", err)
	defer stmtIns.Close()
	_, err = stmtIns.Exec(appId, featureId)
	checkError("app_feature", err)
}
func checkError(table string, err error) {
	if err != nil {
		msg := err.Error()
		if !strings.Contains(msg, "Duplicate") {
			log.Fatalf("Failed to insert into table %s. %s\n", table, msg)
		}
	}
}
