package server

import (
	"database/sql"
	"log"

	"github.com/dimasyanu/simple-dictionary/models"
)

type Dictionary struct {
	db *sql.DB
}

const dbFile string = "store.db"
const create string = `
	CREATE TABLE IF NOT EXISTS dictionary (
		id TEXT NOT NULL PRIMARY KEY,
		term TEXT NOT NULL,
		description TEXT NOT NULL
	);`

func NewDictionary() (*Dictionary, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(create); err != nil {
		return nil, err
	}
	return &Dictionary{
		db: db,
	}, nil
}

func (c *Dictionary) GetItems() ([]models.DictionaryItem, error) {
	rows, err := c.db.Query("SELECT * FROM dictionary ORDER BY term")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []models.DictionaryItem{}
	for rows.Next() {
		var item = models.DictionaryItem{}
		err = rows.Scan(&item.Id, &item.Term, &item.Description)
		if err != nil {
			return nil, err
		}
		results = append(results, item)
	}

	return results, nil
}

func (c *Dictionary) Insert(item models.DictionaryItem) bool {
	_, err := c.db.Exec("INSERT INTO dictionary VALUES(?,?,?);", item.Id, item.Term, item.Description)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (c *Dictionary) Update(item models.DictionaryItem) bool {
	_, err := c.db.Exec("UPDATE dictionary SET term = ?, description = ? WHERE id = ?;", item.Term, item.Description, item.Id)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
