package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/lib/pq"
)

type ItemsRouter struct {
	mux *http.ServeMux
	s   *ItemService
}

func NewItemsRouter(mux *http.ServeMux, db *sql.DB) *ItemsRouter {
	s := NewItemService(db)
	return &ItemsRouter{mux: mux, s: s}
}

func (router *ItemsRouter) InitRoutes() {
	router.mux.HandleFunc("POST /items/", router.createItem)
	router.mux.HandleFunc("GET /items/", router.getAllItems)
	router.mux.HandleFunc("GET /items/{id}", router.getItem)
	router.mux.HandleFunc("PUT /items/{id}", router.updateItem)
	router.mux.HandleFunc("DELETE /items/{id}", router.deleteItem)
}

func (router *ItemsRouter) createItem(w http.ResponseWriter, r *http.Request) {

	var item Item

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err = router.s.createItem(item)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == "23505" {
			http.Error(w, "Item already exists", http.StatusConflict)
			return
		}
		http.Error(w, "Error creating item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

func (router *ItemsRouter) getAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := router.s.getAllItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (router *ItemsRouter) getItem(w http.ResponseWriter, r *http.Request) {

	var item Item

	pathParts := strings.Split(r.URL.Path, "/")
	id, err := strconv.ParseInt(pathParts[2], 10, 64)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	item.Id = id

	item, err = router.s.getItem(item)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Item not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (router *ItemsRouter) updateItem(w http.ResponseWriter, r *http.Request) {

	var item Item
	var newItem Item

	pathParts := strings.Split(r.URL.Path, "/")
	id, err := strconv.ParseInt(pathParts[2], 10, 64)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item.Id = id

	item, err = router.s.updateItem(item, newItem)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Item not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (router *ItemsRouter) deleteItem(w http.ResponseWriter, r *http.Request) {

	var item Item

	pathParts := strings.Split(r.URL.Path, "/")
	id, err := strconv.ParseInt(pathParts[2], 10, 64)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	item.Id = id

	item, err = router.s.deleteItem(item)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Item not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
