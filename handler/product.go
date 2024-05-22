package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"tugas-golang/database"
	"tugas-golang/model"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var product model.Product
	err = json.Unmarshal(payload, &product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request gagal"))
		return
	}

	db := database.ConnectDatabase()
	result := db.Create(&product)
	if result.Error != nil {
		log.Println(result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("request gagal"))
		return
	}

	if result.RowsAffected > 0 {
		response := map[string]interface{}{
			"message":     "Data berhasil disimpan",
			"status_code": http.StatusCreated,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDatabase()

	var products []model.Product
	result := db.Find(&products)
	if result.Error != nil {
		log.Println(result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to fetch products"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
