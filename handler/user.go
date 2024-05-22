package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"tugas-golang/database"
	"tugas-golang/model"
)

func Storeuser(w http.ResponseWriter, r *http.Request) {
	payload, _ := io.ReadAll(r.Body)
	data := []byte(payload)

	user := new(model.User)
	err := json.Unmarshal(data, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request gagal"))
	}

	db := database.ConnectDatabase()
	result := db.Create(user)
	if result.Error != nil {
		log.Fatal(result.Error)
		w.WriteHeader(http.StatusBadRequest)
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDatabase()

	var user []model.User

	if err := db.Find(&user).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
	}

	output, err := json.MarshalIndent(user, "", " ")
	if nil != err {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}
