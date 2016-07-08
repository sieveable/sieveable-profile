package dbwriter

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
)

// Insert the parsed response into MySQL db
func Insert(res Response) (err error) {
	db, err := sql.Open("mysql", os.Getenv("USER")+":"+os.Getenv("PW")+"@/"+os.Getenv("DB"))
	if err != nil {
		return fmt.Errorf("Failed to get a handle for the database %s. %v\n",
			os.Getenv("DB"), err.Error())
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		return fmt.Errorf("Failed to connect to DB. Make sure that the required "+
			"environment variables are set. %v\n", err.Error())
	}
	fmt.Println("Inserting new category", res.Category.Name)
	categoryId, err := insertCategory(db, &res.Category)
	if err != nil {
		return err
	}
	fmt.Println("Inserting new feature", res.Feature.Name)
	featureId, err := insertFeature(db, &res.Feature, &categoryId)
	if err != nil {
		return err
	}
	for _, app := range res.Apps {
		fmt.Printf("Inserting features for %s\n", app.Id)
		if err := insertApp(db, &app); err != nil {
			return err
		} else if err := insertAppFeature(db, app.Id, featureId); err != nil {
			return err
		}
	}
	return nil
}
func insertCategory(db *sql.DB, category *CategoryType) (categoryId int64, err error) {
	stmtIns, err := db.Prepare("INSERT INTO category SET name=?,type=?,description=?")
	if err := checkError("category", err); err != nil {
		return -1, err
	}
	defer stmtIns.Close()
	execResult, err := stmtIns.Exec(category.Name, category.Type, category.Description)
	if err := checkError("category", err); err != nil {
		return -1, err
	}
	if execResult == nil {
		if categoryId, err = getIdByName(db, "category", category.Name); err != nil {
			return -1, err
		}
	} else {
		categoryId, err = execResult.LastInsertId()
		if err != nil {
			return -1, fmt.Errorf("Failed to get category id\n")
		}
	}
	return categoryId, nil
}
func getIdByName(db *sql.DB, table string, name string) (id int64, err error) {
	if table != "category" && table != "feature" {
		return -1, fmt.Errorf("Cannot get id for an unknown table.")
	}
	stmtOut, err := db.Prepare("SELECT id FROM " + table + " WHERE name =?")
	if err != nil {
		return -1, fmt.Errorf("Failed to find "+table+" id for %s\n", name)
	}
	defer stmtOut.Close()
	err = stmtOut.QueryRow(name).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("Failed to find "+table+" id for %s\n", name)
	}
	return id, nil
}
func insertFeature(db *sql.DB, feature *FeatureType, categoryId *int64) (featureId int64, err error) {
	stmtIns, err := db.Prepare("INSERT INTO feature SET " +
		"name=?,description=?,sieveable_query=?,category_id=?")
	if err := checkError("feature", err); err != nil {
		return -1, err
	}
	defer stmtIns.Close()
	execResult, err := stmtIns.Exec(feature.Name, feature.Description,
		feature.SieveableQuery, categoryId)
	if err = checkError("feature", err); err != nil {
		return -1, err
	}
	if execResult == nil {
		if featureId, err = getIdByName(db, "feature", feature.Name); err != nil {
			return -1, err
		}
	} else {
		featureId, err = execResult.LastInsertId()
		if err != nil {
			return -1, fmt.Errorf("Failed to get feature id")
		}
	}
	return featureId, nil
}
func insertApp(db *sql.DB, app *AppType) (err error) {
	stmtIns, err := db.Prepare("INSERT INTO app SET id=?, package_name=?," +
		"version_code=?,version_name=?,downloads=?,ratings=?,release_date=?")
	if err = checkError("app", err); err != nil {
		return err
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(app.Id, app.PackageName, app.VersionCode,
		app.VersionName, app.Listing.Downloads, app.Listing.Ratings,
		app.Listing.ReleaseDate.Format("2006-01-02"))
	if err = checkError("app", err); err != nil {
		return err
	}
	return nil
}
func insertAppFeature(db *sql.DB, appId string, featureId int64) (err error) {
	stmtIns, err := db.Prepare("INSERT INTO app_feature VALUES(?,?)")
	if err = checkError("app_feature", err); err != nil {
		return err
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(appId, featureId)
	if err = checkError("app_feature", err); err != nil {
		return err
	}
	return nil
}
func checkError(table string, e error) (err error) {
	if e != nil {
		msg := e.Error()
		if !strings.Contains(msg, "Duplicate") {
			return fmt.Errorf("Failed to insert into table %s. %s\n", table, msg)
		}
	}
	return nil
}
