package main

import (
	"database/sql"
	"time"
)

type ItemService struct {
	db *sql.DB
}

func NewItemService(db *sql.DB) *ItemService {
	return &ItemService{db: db}
}

func (i *ItemService) createItem(item Item) (Item, error) {
	query := `
	INSERT INTO items (name)
	VALUES ($1)
	RETURNING id, name, date
	`

	var dbItem Item

	err := i.db.QueryRow(query, item.Name).Scan(&dbItem.Id, &dbItem.Name, &dbItem.Date)
	if err != nil {
		return Item{}, err
	}

	return dbItem, nil
}

func (i *ItemService) getAllItems() ([]Item, error) {
	query := `
	SELECT * FROM items
	`

	data := []Item{}

	rows, err := i.db.Query(query)
	if err != nil {
		return []Item{}, err
	}

	defer rows.Close()

	var id int64
	var name string
	var date time.Time

	for rows.Next() {
		err := rows.Scan(&id, &name, &date)
		if err != nil {
			return []Item{}, err
		}

		data = append(data, Item{Id: id, Name: name, Date: date})
	}

	return data, nil
}

func (i *ItemService) getItem(item Item) (Item, error) {
	query := `
	SELECT * FROM items WHERE id = $1
	`

	var dbItem Item

	err := i.db.QueryRow(query, item.Id).Scan(&dbItem.Id, &dbItem.Name, &dbItem.Date)
	if err != nil {
		return Item{}, err
	}

	return dbItem, nil
}

func (i *ItemService) updateItem(item Item, newItem Item) (Item, error) {
	query := `
	UPDATE items
	SET name = $2, date = NOW()
	WHERE id = $1
	RETURNING id, name, date
	`

	var updatedItem Item

	err := i.db.QueryRow(query, item.Id, newItem.Name).Scan(&updatedItem.Id, &updatedItem.Name, &updatedItem.Date)
	if err != nil {
		return Item{}, err
	}

	return updatedItem, nil
}

func (i *ItemService) deleteItem(item Item) (Item, error) {
	query := `
	DELETE FROM items 
	WHERE id = $1
	RETURNING id, name, date
	`

	var deletedItem Item

	err := i.db.QueryRow(query, item.Id).Scan(&deletedItem.Id, &deletedItem.Name, &deletedItem.Name)
	if err != nil {
		return Item{}, err
	}

	return deletedItem, nil
}
