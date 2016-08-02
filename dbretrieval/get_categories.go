package dbretrieval

import (
	"database/sql"
)

type categoryResult struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func GetCategoriesByType(db *sql.DB, typeName string) ([]categoryResult, error) {
	rows, err := db.Query("SELECT * FROM category WHERE type=?", typeName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []categoryResult{}
	for rows.Next() {
		category := categoryResult{}
		err := rows.Scan(&category.Id, &category.Name, &category.Type,
			&category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
