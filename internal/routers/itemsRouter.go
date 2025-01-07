package routers

import (
	"database/sql"
	"encoding/json"
	"log"
	"main/internal/middlewares"
	"main/internal/services"
	"main/internal/types"
	"net/http"
	"strconv"
	"strings"

	"github.com/lib/pq"
)

type ItemsRouter struct {
	mux *http.ServeMux
	s   *services.ItemService
}

type Item = types.Item

func NewItemsRouter(mux *http.ServeMux, db *sql.DB) *ItemsRouter {
	s := services.NewItemService(db)
	err := s.CreateItemsTable()
	if err != nil {
		log.Fatal("Error creating items table:", err)
	}
	return &ItemsRouter{mux: mux, s: s}
}

func (router *ItemsRouter) InitRoutes() {
	router.mux.HandleFunc("POST /items/", middlewares.AuthMiddleware(router.createItem))
	router.mux.HandleFunc("GET /items/", router.getAllItems)
	router.mux.HandleFunc("GET /items/{id}", router.getItem)
	router.mux.HandleFunc("PUT /items/{id}", middlewares.AuthMiddleware(router.updateItem))
	router.mux.HandleFunc("DELETE /items/{id}", middlewares.AuthMiddleware(router.deleteItem))
}

func (router *ItemsRouter) createItem(w http.ResponseWriter, r *http.Request) {

	var item Item

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err = router.s.CreateItem(item)
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
	items, err := router.s.GetAllItems()
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

	item, err = router.s.GetItem(item)
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

	item, err = router.s.UpdateItem(item, newItem)
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

	item, err = router.s.DeleteItem(item)
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
