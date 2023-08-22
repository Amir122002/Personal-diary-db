package handlers

import (
	"diary_db/internal/database"
	"diary_db/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

var db, _ = database.RunDb()

func Create(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var requestBody map[string]interface{}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	textJSON, ok := requestBody["text"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Field 'text' not found in JSON")
		return
	}

	text, ok := textJSON.(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Field 'text' is not a string")
		return
	}
	db.Exec("insert into DiaryEntry(text) values ($1)", text)
	w.WriteHeader(http.StatusOK)
}

func Read(w http.ResponseWriter, r *http.Request) {
	var diary []models.DiaryEntry
	db.Raw("SELECT * FROM DiaryEntry where active=true").Scan(&diary)
	jsonBytes, err := json.Marshal(diary)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var requestBody map[string]interface{}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	textJSON, exists := requestBody["text"]
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Field 'text' not found in JSON")
		return
	}

	text, ok := textJSON.(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Field 'text' is not a string")
		return
	}
	//query := "UPDATE DiaryEntry set text=$1, update_at=current_timestamp where id=$2"
	db.Exec("UPDATE DiaryEntry set text=$1, update_at=current_timestamp where id=$2", text, id)
	w.WriteHeader(http.StatusOK)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	db.Exec("update DiaryEntry set active=false, delete_at=current_timestamp where id=$1", id)
	w.WriteHeader(http.StatusOK)
}
